using System.IO.Pipes;
using System.Text;

namespace AppPulse.WinListener.Services;

public class IpcService : IDisposable
{
    private const string PipeName = "AppPulsePipe";
    private NamedPipeServerStream? _pipeServer;
    private StreamWriter? _streamWriter;
    private readonly CancellationTokenSource _cancellationTokenSource = new();

    public IpcService()
    {
    }

    public async Task StartServerAsync(CancellationToken cancellationToken)
    {
        try
        {
            _pipeServer = new NamedPipeServerStream(PipeName, PipeDirection.Out, 1, PipeTransmissionMode.Byte, PipeOptions.Asynchronous);

            Console.WriteLine($"[{DateTime.Now:HH:mm:ss}] IPC 伺服器已創建命名管道: {PipeName}");
            Console.WriteLine($"[{DateTime.Now:HH:mm:ss}] IPC 伺服器正在等待連線...");

            await _pipeServer.WaitForConnectionAsync(cancellationToken);

            _streamWriter = new StreamWriter(_pipeServer, Encoding.UTF8) { AutoFlush = true };
            Console.WriteLine($"[{DateTime.Now:HH:mm:ss}] IPC 客戶端已連線。");
        }
        catch (OperationCanceledException)
        {
            Console.WriteLine($"[{DateTime.Now:HH:mm:ss}] IPC 伺服器等待連線已取消。");
        }
        catch (Exception ex)
        {
            Console.WriteLine($"[{DateTime.Now:HH:mm:ss}] IPC 伺服器啟動失敗: {ex.Message}");
        }
    }

    public async Task SendMessageAsync(string message)
    {
        if (_pipeServer != null && _pipeServer.IsConnected && _streamWriter != null)
        {
            try
            {
                await _streamWriter.WriteLineAsync(message);
            }
            catch (IOException ex)
            {
                Console.WriteLine($"[{DateTime.Now:HH:mm:ss}] 寫入 IPC 管道時發生錯誤: {ex.Message}");
                // Handle client disconnection if necessary
            }
        }
    }

    public void Dispose()
    {
        _cancellationTokenSource.Cancel();
        _streamWriter?.Dispose();
        _pipeServer?.Dispose();
        _cancellationTokenSource.Dispose();
        GC.SuppressFinalize(this);
    }
}