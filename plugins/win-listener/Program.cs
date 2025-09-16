using AppPulse.WinListener.Services;
using System.Text.Json;
using System.Text.Json.Serialization;
using Windows.Media.Control;
using System.IO;
using System.Text;

namespace AppPulse.WinListener;

class Program
{
    private static readonly CancellationTokenSource _cancellationTokenSource = new();
    private static ProcessTraceService? _processTraceService;
    private static MediaTraceService? _mediaTraceService;
    private static IpcService? _ipcService;

    private static string _lastWindowTitle = string.Empty;
    private static string _lastProcessName = string.Empty;
    private static string _lastTrackTitle = string.Empty;
    private static string _lastArtist = string.Empty;
    private static GlobalSystemMediaTransportControlsSessionPlaybackStatus _lastStatus = GlobalSystemMediaTransportControlsSessionPlaybackStatus.Closed;


    static async Task Main(string[] args)
    {
        // Force UTF-8 encoding for stdout and stderr to ensure proper communication with the Go host.
        Console.OutputEncoding = System.Text.Encoding.UTF8;
        // `Console.ErrorEncoding` is not available on all targets. Replace by setting a UTF-8 StreamWriter for stderr.
        var stderrWriter = new StreamWriter(Console.OpenStandardError(), Encoding.UTF8)
        {
            AutoFlush = true
        };
        Console.SetError(stderrWriter);

        Console.WriteLine("AppPulse - 前台窗口與音樂播放監聽服務");
        Console.WriteLine("按 Ctrl+C 退出程序");
        Console.WriteLine();

        var cancellationToken = _cancellationTokenSource.Token;

        Console.CancelKeyPress += (sender, e) =>
        {
            e.Cancel = true;
            _cancellationTokenSource.Cancel();
        };

        try
        {
            _processTraceService = new ProcessTraceService();
            _mediaTraceService = await MediaTraceService.CreateAsync();
            _ipcService = new IpcService();

            _processTraceService.OnFrontWindowChanged += OnFrontWindowChanged;
            _mediaTraceService.OnMediaPlaybackChanged += OnMediaPlaybackChanged;

            // Run services in the background
            Task.Run(() => _processTraceService.StartAsync(cancellationToken), cancellationToken);
            Task.Run(() => _ipcService.StartServerAsync(cancellationToken), cancellationToken);

            // Keep the main thread alive
            cancellationToken.WaitHandle.WaitOne();
        }
        catch (OperationCanceledException)
        {
            Console.WriteLine("\n程序正在退出...");
        }
        catch (Exception ex)
        {
            Console.Error.WriteLine($"程序發生錯誤: {ex.Message}");
            Console.Error.WriteLine(ex.StackTrace);
        }
        finally
        {
            _processTraceService?.Dispose();
            _mediaTraceService?.Dispose();
            _ipcService?.Dispose();
            _cancellationTokenSource.Dispose();
        }
    }

    private static void OnFrontWindowChanged(string windowTitle, string displayName, string originalProcessName)
    {
        if (windowTitle == _lastWindowTitle && originalProcessName == _lastProcessName) return;

        _lastWindowTitle = windowTitle;
        _lastProcessName = originalProcessName;

        Console.WriteLine($"[{DateTime.Now:HH:mm:ss}] 前台窗口: {originalProcessName} # {displayName} # {windowTitle}");

        var windowEvent = new WindowEvent
        {
            Time = DateTime.Now.ToString("HH:mm:ss"),
            Process = originalProcessName,
            App = displayName,
            Title = windowTitle
        };

        string json = JsonSerializer.Serialize(windowEvent, JsonOptions.Default);
        _ipcService?.SendMessageAsync(json);
    }

    private static void OnMediaPlaybackChanged(string title, string artist, GlobalSystemMediaTransportControlsSessionPlaybackStatus status)
    {
        if (title == _lastTrackTitle && artist == _lastArtist && status == _lastStatus) return;

        _lastTrackTitle = title;
        _lastArtist = artist;
        _lastStatus = status;

        var statusText = GetPlaybackStatusText(status);
        Console.WriteLine($"[{DateTime.Now:HH:mm:ss}] 音樂播放: {statusText}");
        if (status is GlobalSystemMediaTransportControlsSessionPlaybackStatus.Playing or GlobalSystemMediaTransportControlsSessionPlaybackStatus.Paused)
        {
            Console.WriteLine($"  曲名: {title}");
            Console.WriteLine($"  演出者: {artist}");
        }
        Console.WriteLine();

        if (status is GlobalSystemMediaTransportControlsSessionPlaybackStatus.Playing or GlobalSystemMediaTransportControlsSessionPlaybackStatus.Paused)
        {
            var musicEvent = new MusicEvent
            {
                Time = DateTime.Now.ToString("HH:mm:ss"),
                Status = status == GlobalSystemMediaTransportControlsSessionPlaybackStatus.Playing ? "playing" : "paused",
                Title = title,
                Artist = artist
            };
            string json = JsonSerializer.Serialize(musicEvent, JsonOptions.Default);
            _ipcService?.SendMessageAsync(json);
        }
    }

    private static string GetPlaybackStatusText(GlobalSystemMediaTransportControlsSessionPlaybackStatus status)
    {
        return status switch
        {
            GlobalSystemMediaTransportControlsSessionPlaybackStatus.Playing => "▶ 播放中",
            GlobalSystemMediaTransportControlsSessionPlaybackStatus.Paused => "⏸ 暫停",
            GlobalSystemMediaTransportControlsSessionPlaybackStatus.Stopped => "⏹ 停止",
            GlobalSystemMediaTransportControlsSessionPlaybackStatus.Changing => "🔄 切換中",
            GlobalSystemMediaTransportControlsSessionPlaybackStatus.Closed => "❌ 關閉",
            _ => "❓ 未知狀態"
        };
    }
}

// Data models for JSON serialization
public abstract class BaseEvent
{
    [JsonPropertyName("time")]
    public required string Time { get; set; }
    [JsonPropertyName("type")]
    public abstract string Type { get; }
}

public class MusicEvent : BaseEvent
{
    public override string Type => "music";
    [JsonPropertyName("status")]
    public required string Status { get; set; }
    [JsonPropertyName("title")]
    public required string Title { get; set; }
    [JsonPropertyName("artist")]
    public required string Artist { get; set; }
}

public class WindowEvent : BaseEvent
{
    public override string Type => "window";
    [JsonPropertyName("process")]
    public required string Process { get; set; }
    [JsonPropertyName("app")]
    public required string App { get; set; }
    [JsonPropertyName("title")]
    public required string Title { get; set; }
}

public static class JsonOptions
{
    public static readonly JsonSerializerOptions Default = new()
    {
        PropertyNamingPolicy = JsonNamingPolicy.CamelCase,
        WriteIndented = false
    };
}
