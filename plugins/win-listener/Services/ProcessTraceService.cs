using System.Diagnostics;
using static Windows.Win32.PInvoke;
using Windows.Win32.UI.Accessibility;
using Windows.Win32.Foundation;
using Windows.Win32.UI.WindowsAndMessaging;

namespace AppPulse.WinListener.Services;

public class ProcessTraceService : IDisposable
{
    private const uint EventObjectNameChange = 0x800C;
    private const int ObjIdWindow = 0;
    private static readonly TimeSpan NameChangeDebounce = TimeSpan.FromMilliseconds(400);

    private HWINEVENTHOOK _hookHandle;
    private HWINEVENTHOOK _nameHookHandle;
    private readonly System.Threading.Timer _nameChangeTimer;
    private HWND _pendingNameHwnd;
    private readonly WINEVENTPROC _hookDelegate;
    private bool _disposed = false;
    private HWND _lastHwnd;
    private readonly Thread _eventThread;
    private uint _eventThreadId;

    public delegate void FrontWindowChangeHandler(string windowTitle, string displayName, string originalProcessName);
    public event FrontWindowChangeHandler? OnFrontWindowChanged;

    public ProcessTraceService()
    {
        _hookDelegate = new WINEVENTPROC(WinEventProc);
        _nameChangeTimer = new System.Threading.Timer(_ => FlushNameChange(), null, Timeout.Infinite, Timeout.Infinite);
        _eventThread = new Thread(() =>
        {
            // Store the thread ID for PostThreadMessage
            _eventThreadId = GetCurrentThreadId();

            // Set the hook on this thread
            _hookHandle = SetWinEventHook(
                EVENT_SYSTEM_FOREGROUND,
                EVENT_SYSTEM_FOREGROUND,
                HMODULE.Null,
                _hookDelegate,
                0, 0,
                WINEVENT_OUTOFCONTEXT);

            // 前景視窗標題變化（例如 Edge 切換分頁）不會觸發前景事件，另外監聽。
            _nameHookHandle = SetWinEventHook(
                EventObjectNameChange,
                EventObjectNameChange,
                HMODULE.Null,
                _hookDelegate,
                0, 0,
                WINEVENT_OUTOFCONTEXT);

            // Initial check
            CheckCurrentWindow();

            // Run the message loop
            while (GetMessage(out var msg, HWND.Null, 0, 0))
            {
                TranslateMessage(in msg);
                DispatchMessage(in msg);
            }
        });
        _eventThread.SetApartmentState(ApartmentState.STA);
        _eventThread.IsBackground = true;
    }

    private void CheckCurrentWindow()
    {
        var hwnd = GetForegroundWindow();
        if (hwnd.Value != IntPtr.Zero)
        {
            InvokeHandlers(hwnd);
        }
    }

    private void WinEventProc(HWINEVENTHOOK hWinEventHook, uint eventType, HWND hwnd, int idObject, int idChild, uint dwEventThread, uint dwmsEventTime)
    {
        if (_disposed) return;

        try
        {
            if (eventType == EventObjectNameChange)
            {
                if (idObject != ObjIdWindow || idChild != 0 || hwnd.Value == IntPtr.Zero) return;
                if (hwnd != GetForegroundWindow()) return;
                _pendingNameHwnd = hwnd;
                _nameChangeTimer.Change(NameChangeDebounce, Timeout.InfiniteTimeSpan);
                return;
            }

            if (hwnd.Value == IntPtr.Zero || hwnd == _lastHwnd) return;
            _lastHwnd = hwnd;
            InvokeHandlers(hwnd);
        }
        catch (Exception ex)
        {
            Console.WriteLine($"[{DateTime.Now:HH:mm:ss}] WinEventProc 錯誤: {ex.Message}");
        }
    }

    private void FlushNameChange()
    {
        if (_disposed) return;
        try
        {
            var hwnd = _pendingNameHwnd;
            if (hwnd.Value != IntPtr.Zero && hwnd == GetForegroundWindow())
            {
                InvokeHandlers(hwnd);
            }
        }
        catch (Exception ex)
        {
            Console.WriteLine($"[{DateTime.Now:HH:mm:ss}] 標題變更處理錯誤: {ex.Message}");
        }
    }

    private void InvokeHandlers(HWND hwnd)
    {
        var (title, displayName, originalProcessName) = GetWindowDetails(hwnd);
        OnFrontWindowChanged?.Invoke(title, displayName, originalProcessName);
    }

    private unsafe (string windowTitle, string displayName, string originalProcessName) GetWindowDetails(HWND hWnd)
    {
        try
        {
            uint processId = 0;
            GetWindowThreadProcessId(hWnd, &processId);
            if (processId == 0) return ("未知標題", "未知進程", "未知進程");

            using var process = Process.GetProcessById((int)processId);
            string originalProcessName = process.ProcessName;
            string displayName = originalProcessName;

            try
            {
                // Try to get a friendly name from the file version info (ProductName or FileDescription)
                var mainModule = process.MainModule;
                if (mainModule is not null)
                {
                    var fvi = mainModule.FileVersionInfo;
                    string? friendly = null;
                    if (!string.IsNullOrWhiteSpace(fvi.ProductName))
                    {
                        friendly = fvi.ProductName;
                    }
                    else if (!string.IsNullOrWhiteSpace(fvi.FileDescription))
                    {
                        friendly = fvi.FileDescription;
                    }

                    if (!string.IsNullOrWhiteSpace(friendly))
                    {
                        // Keep both: "Friendly Name (processname)" so system names remain identifiable
                        displayName = friendly;
                    }
                }
            }
            catch
            {
                // Accessing MainModule/FileVersionInfo may fail for system or protected processes.
                // Keep original process name as displayName.
            }

            const int nMaxCount = 256;
            var titleBuffer = new char[nMaxCount];
            int length;
            fixed (char* pTitleBuffer = titleBuffer)
            {
                length = GetWindowText(hWnd, pTitleBuffer, nMaxCount);
            }

            var windowTitle = length > 0 ? new string(titleBuffer, 0, length) : "無標題";

            return (windowTitle, displayName, originalProcessName);
        }
        catch
        {
            return ("未知標題", "未知進程", "未知進程");
        }
    }

    public Task StartAsync(CancellationToken cancellationToken)
    {
        Console.WriteLine("前台窗口監聽服務已啟動");
        _eventThread.Start();
        cancellationToken.Register(Dispose);
        return Task.CompletedTask;
    }

    public void Dispose()
    {
        if (!_disposed)
        {
            _disposed = true;
            if (_hookHandle.Value != IntPtr.Zero)
            {
                UnhookWinEvent(_hookHandle);
            }
            if (_nameHookHandle.Value != IntPtr.Zero)
            {
                UnhookWinEvent(_nameHookHandle);
            }
            _nameChangeTimer.Dispose();
            // Safely stop the message loop thread
            if (_eventThread.IsAlive)
            {
                PostThreadMessage(_eventThreadId, WM_QUIT, default, default);
                _eventThread.Join(500); // Wait for the thread to exit
            }
            GC.SuppressFinalize(this);
        }
    }
}
