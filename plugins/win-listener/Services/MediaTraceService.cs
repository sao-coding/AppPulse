using Windows.Media.Control;

namespace AppPulse.WinListener.Services;

public class MediaTraceService : IDisposable
{
    public delegate void MediaPlaybackChangeHandler(string title, string artist, GlobalSystemMediaTransportControlsSessionPlaybackStatus status);

    public event MediaPlaybackChangeHandler? OnMediaPlaybackChanged;

    private GlobalSystemMediaTransportControlsSessionManager? _sessionManager;
    private GlobalSystemMediaTransportControlsSession? _currentSession;
    private bool _disposed = false;
    private CancellationTokenSource? _debounceCts;

    // State tracking to avoid duplicate events
    private string _lastTrackTitle = string.Empty;
    private string _lastArtist = string.Empty;
    public GlobalSystemMediaTransportControlsSessionPlaybackStatus LastStatus { get; private set; } = GlobalSystemMediaTransportControlsSessionPlaybackStatus.Closed;

    private MediaTraceService()
    {
    }

    public static async Task<MediaTraceService> CreateAsync()
    {
        var service = new MediaTraceService();
        await service.InitializeAsync();
        return service;
    }

    private async Task InitializeAsync()
    {
        try
        {
            _sessionManager = await GlobalSystemMediaTransportControlsSessionManager.RequestAsync();
            if (_sessionManager != null)
            {
                _sessionManager.SessionsChanged += OnSessionsChanged;
                // Manually trigger the event to get the current state
                OnSessionsChanged(_sessionManager, null);
            }
        }
        catch (Exception ex)
        {
            Console.WriteLine($"[{DateTime.Now:HH:mm:ss}] 媒體會話初始化錯誤: {ex.Message}");
        }
    }

    private async void OnSessionsChanged(GlobalSystemMediaTransportControlsSessionManager sender, SessionsChangedEventArgs? args)
    {
        this._debounceCts?.Cancel();
        this._debounceCts = new CancellationTokenSource();

        try
        {
            await Task.Delay(250, this._debounceCts.Token);
            ProcessSessionChange(sender);
        }
        catch (TaskCanceledException)
        {
            // This is expected if a new event comes in quickly.
        }
    }

    private void ProcessSessionChange(GlobalSystemMediaTransportControlsSessionManager sender)
    {
        if (_currentSession != null)
        {
            _currentSession.MediaPropertiesChanged -= OnMediaPropertiesChanged;
            _currentSession.PlaybackInfoChanged -= OnPlaybackInfoChanged;
        }

        var sessions = sender.GetSessions();
        var newSession = sessions.FirstOrDefault(s => s.GetPlaybackInfo().PlaybackStatus == GlobalSystemMediaTransportControlsSessionPlaybackStatus.Playing)
                         ?? sessions.FirstOrDefault(s => s.GetPlaybackInfo().PlaybackStatus == GlobalSystemMediaTransportControlsSessionPlaybackStatus.Paused);

        _currentSession = newSession;

        if (_currentSession != null)
        {
            _currentSession.MediaPropertiesChanged += OnMediaPropertiesChanged;
            _currentSession.PlaybackInfoChanged += OnPlaybackInfoChanged;
            UpdateMediaState(_currentSession);
        }
        else
        {
            UpdateMediaState(null);
        }
    }

    private void OnPlaybackInfoChanged(GlobalSystemMediaTransportControlsSession sender, PlaybackInfoChangedEventArgs args)
    {
        UpdateMediaState(sender);
    }

    private void OnMediaPropertiesChanged(GlobalSystemMediaTransportControlsSession sender, MediaPropertiesChangedEventArgs args)
    {
        UpdateMediaState(sender);
    }

    private async void UpdateMediaState(GlobalSystemMediaTransportControlsSession? session)
    {
        if (session == null)
        {
            if (LastStatus != GlobalSystemMediaTransportControlsSessionPlaybackStatus.Closed)
            {
                LastStatus = GlobalSystemMediaTransportControlsSessionPlaybackStatus.Closed;
                _lastTrackTitle = string.Empty;
                _lastArtist = string.Empty;
                OnMediaPlaybackChanged?.Invoke(string.Empty, string.Empty, LastStatus);
            }
            return;
        }

        try
        {
            var mediaProperties = await session.TryGetMediaPropertiesAsync();
            var playbackInfo = session.GetPlaybackInfo();

            var currentTitle = mediaProperties?.Title ?? string.Empty;
            var currentArtist = mediaProperties?.Artist ?? string.Empty;
            var currentStatus = playbackInfo?.PlaybackStatus ?? GlobalSystemMediaTransportControlsSessionPlaybackStatus.Closed;

            // Filter out events that are likely ads or other non-music media
            if ((currentStatus == GlobalSystemMediaTransportControlsSessionPlaybackStatus.Playing ||
                 currentStatus == GlobalSystemMediaTransportControlsSessionPlaybackStatus.Paused) &&
                string.IsNullOrEmpty(currentArtist))
            {
                return;
            }

            if (currentTitle != _lastTrackTitle || currentArtist != _lastArtist || currentStatus != LastStatus)
            {
                _lastTrackTitle = currentTitle;
                _lastArtist = currentArtist;
                LastStatus = currentStatus;

                OnMediaPlaybackChanged?.Invoke(currentTitle, currentArtist, currentStatus);
            }
        }
        catch (Exception ex)
        {
            Console.WriteLine($"[{DateTime.Now:HH:mm:ss}] 更新媒體狀態時發生錯誤: {ex.Message}");
        }
    }

    public void Dispose()
    {
        if (!_disposed)
        {
            if (_currentSession != null)
            {
                _currentSession.MediaPropertiesChanged -= OnMediaPropertiesChanged;
                _currentSession.PlaybackInfoChanged -= OnPlaybackInfoChanged;
            }

            if (_sessionManager != null)
            {
                _sessionManager.SessionsChanged -= OnSessionsChanged;
            }
            _disposed = true;
            GC.SuppressFinalize(this);
        }
    }
}