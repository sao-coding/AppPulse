using Microsoft.Win32;
using System.Runtime.InteropServices;

namespace AppPulse.WinListener.Services;

/// <summary>
/// 監聽工作階段鎖定/解鎖、睡眠/喚醒，並定期回報距離上次鍵鼠輸入的秒數。
/// </summary>
public class SystemStateService : IDisposable
{
    private static readonly TimeSpan InputPollInterval = TimeSpan.FromSeconds(5);

    private readonly System.Threading.Timer _inputTimer;
    private bool _disposed;

    /// <summary>state: locked / unlocked / suspend / resume</summary>
    public event Action<string>? OnSessionState;

    /// <summary>距離上次鍵鼠輸入的秒數。</summary>
    public event Action<int>? OnInputIdle;

    public SystemStateService()
    {
        _inputTimer = new System.Threading.Timer(_ => PollInput(), null, Timeout.Infinite, Timeout.Infinite);
    }

    public void Start()
    {
        SystemEvents.SessionSwitch += HandleSessionSwitch;
        SystemEvents.PowerModeChanged += HandlePowerMode;
        _inputTimer.Change(TimeSpan.Zero, InputPollInterval);
    }

    private void HandleSessionSwitch(object? sender, SessionSwitchEventArgs e)
    {
        switch (e.Reason)
        {
            case SessionSwitchReason.SessionLock:
            case SessionSwitchReason.ConsoleDisconnect:
            case SessionSwitchReason.RemoteDisconnect:
                OnSessionState?.Invoke("locked");
                break;
            case SessionSwitchReason.SessionUnlock:
            case SessionSwitchReason.ConsoleConnect:
            case SessionSwitchReason.RemoteConnect:
                OnSessionState?.Invoke("unlocked");
                break;
        }
    }

    private void HandlePowerMode(object? sender, PowerModeChangedEventArgs e)
    {
        if (e.Mode == PowerModes.Suspend) OnSessionState?.Invoke("suspend");
        else if (e.Mode == PowerModes.Resume) OnSessionState?.Invoke("resume");
    }

    [StructLayout(LayoutKind.Sequential)]
    private struct LASTINPUTINFO
    {
        public uint cbSize;
        public uint dwTime;
    }

    [DllImport("user32.dll")]
    private static extern bool GetLastInputInfo(ref LASTINPUTINFO plii);

    private void PollInput()
    {
        if (_disposed) return;
        try
        {
            var info = new LASTINPUTINFO { cbSize = (uint)Marshal.SizeOf<LASTINPUTINFO>() };
            if (!GetLastInputInfo(ref info)) return;
            // 兩者皆為 32 位元 tick，相減自然處理溢位。
            uint idleMs = unchecked((uint)Environment.TickCount - info.dwTime);
            OnInputIdle?.Invoke((int)(idleMs / 1000));
        }
        catch (Exception ex)
        {
            Console.WriteLine($"[{DateTime.Now:HH:mm:ss}] 讀取輸入閒置時間失敗: {ex.Message}");
        }
    }

    public void Dispose()
    {
        if (_disposed) return;
        _disposed = true;
        _inputTimer.Dispose();
        SystemEvents.SessionSwitch -= HandleSessionSwitch;
        SystemEvents.PowerModeChanged -= HandlePowerMode;
        GC.SuppressFinalize(this);
    }
}
