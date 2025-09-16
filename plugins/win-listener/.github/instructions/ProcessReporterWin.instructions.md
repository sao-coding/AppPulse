---
applyTo: "**/*.cs"
---

Repository: chingcdesu/shiroprocessreporter
Files analyzed: 45

Estimated tokens: 25.3k

Directory structure:
└── chingcdesu-shiroprocessreporter/
├── README.md
├── ShiroProcessReporter.sln
└── ShiroProcessReporter/
├── ShiroProcessReporter/
│ ├── app.manifest
│ ├── App.xaml
│ ├── App.xaml.cs
│ ├── MainView.xaml
│ ├── MainView.xaml.cs
│ ├── NativeMethods.txt
│ ├── ShiroProcessReporter.csproj
│ ├── Components/
│ │ ├── TrayIconView.xaml
│ │ └── TrayIconView.xaml.cs
│ ├── Converters/
│ │ ├── AppThemeToImageSourceConverter.cs
│ │ ├── LogLevelToColorConverter.cs
│ │ └── ReplaceRulesDictionaryToArrayConverter.cs
│ ├── Extensions/
│ │ └── MemoryLoggerExtension.cs
│ ├── Helper/
│ │ ├── AppLogHelper.cs
│ │ ├── DataContextWrapper.cs
│ │ ├── Debouncer.cs
│ │ ├── GlobalState.cs
│ │ ├── GlobalStateWrapper.cs
│ │ └── Preferences.cs
│ ├── Layouts/
│ │ ├── Navigation.xaml
│ │ └── Navigation.xaml.cs
│ ├── Loggers/
│ │ └── MemoryLogger.cs
│ ├── Models/
│ │ ├── FilterRule.cs
│ │ ├── LogEntry.cs
│ │ └── ReplaceRule.cs
│ ├── Properties/
│ │ └── PublishProfiles/
│ │ ├── win-arm64.pubxml
│ │ ├── win-x64.pubxml
│ │ └── win-x86.pubxml
│ ├── Providers/
│ │ └── MemoryLoggerProvider.cs
│ ├── Resources/
│ │ ├── Images/
│ │ └── TrayIcon/
│ ├── Services/
│ │ ├── MediaTraceService.cs
│ │ ├── ProcessTraceService.cs
│ │ ├── ReportService.cs
│ │ └── TraceWorkerService.cs
│ └── Views/
│ ├── EndpointSettings.xaml
│ ├── EndpointSettings.xaml.cs
│ ├── FilterRuleSettings.xaml
│ ├── FilterRuleSettings.xaml.cs
│ ├── LogViewer.xaml
│ ├── LogViewer.xaml.cs
│ ├── ReplaceRuleSettings.xaml
│ └── ReplaceRuleSettings.xaml.cs
└── ShiroProcessReporter (Package)/
├── Package.appxmanifest
├── ShiroProcessReporter (Package).wapproj
└── Images/

================================================
FILE: README.md
================================================

# ProcessReporterWin

[ProcessReporterWinpy](https://github.com/TNXG/ProcessReporterWinpy)的 C#实现，魔改了部分模块。

### 主要修改部分

- 前台窗口监听使用 Win32 Hook
- 正在播放的音乐信息使用 WinRT 的 MediaSession 监听

### 安装

- 下载安装程序 [Releases](https://github.com/ChingCdesu/ShiroProcessReporter/releases)
- 双击安装应用
  - 如果提示证书问题，下载这个证书文件：[Certificate.cer](https://github.com/ChingCdesu/ShiroProcessReporter/releases/tag/v2.0.0-beta.1)
  - 双击证书文件，选择安装证书
  - 储存位置选 `本地计算机` 下一步
  - 证书储存选择 `将所有的证书都放入下列存储`
  - 点浏览，选择 `受信任的根证书颁发机构`
  - 下一步知道完成安装
  - 重新双击安装应用

================================================
FILE: ShiroProcessReporter.sln
================================================

Microsoft Visual Studio Solution File, Format Version 12.00

# Visual Studio Version 17

VisualStudioVersion = 17.8.34525.116
MinimumVisualStudioVersion = 10.0.40219.1
Project("{C7167F0D-BC9F-4E6E-AFE1-012C56B48DB5}") = "ShiroProcessReporter (Package)", "ShiroProcessReporter\ShiroProcessReporter (Package)\ShiroProcessReporter (Package).wapproj", "{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}"
EndProject
Project("{9A19103F-16F7-4668-BE54-9A1E7A4F7556}") = "ShiroProcessReporter", "ShiroProcessReporter\ShiroProcessReporter\ShiroProcessReporter.csproj", "{A3B50604-0EAA-45F4-9991-39BD878E65A0}"
EndProject
Global
GlobalSection(SolutionConfigurationPlatforms) = preSolution
Debug|Any CPU = Debug|Any CPU
Debug|ARM64 = Debug|ARM64
Debug|x64 = Debug|x64
Debug|x86 = Debug|x86
Release|Any CPU = Release|Any CPU
Release|ARM64 = Release|ARM64
Release|x64 = Release|x64
Release|x86 = Release|x86
EndGlobalSection
GlobalSection(ProjectConfigurationPlatforms) = postSolution
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Debug|Any CPU.ActiveCfg = Debug|x64
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Debug|Any CPU.Build.0 = Debug|x64
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Debug|Any CPU.Deploy.0 = Debug|x64
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Debug|ARM64.ActiveCfg = Debug|ARM64
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Debug|ARM64.Build.0 = Debug|ARM64
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Debug|ARM64.Deploy.0 = Debug|ARM64
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Debug|x64.ActiveCfg = Debug|x64
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Debug|x64.Build.0 = Debug|x64
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Debug|x64.Deploy.0 = Debug|x64
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Debug|x86.ActiveCfg = Debug|x86
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Debug|x86.Build.0 = Debug|x86
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Debug|x86.Deploy.0 = Debug|x86
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Release|Any CPU.ActiveCfg = Release|x64
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Release|Any CPU.Build.0 = Release|x64
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Release|Any CPU.Deploy.0 = Release|x64
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Release|ARM64.ActiveCfg = Release|ARM64
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Release|ARM64.Build.0 = Release|ARM64
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Release|ARM64.Deploy.0 = Release|ARM64
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Release|x64.ActiveCfg = Release|x64
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Release|x64.Build.0 = Release|x64
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Release|x64.Deploy.0 = Release|x64
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Release|x86.ActiveCfg = Release|x86
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Release|x86.Build.0 = Release|x86
{AC822E31-7C2D-4F32-884B-DB0B5FDDA414}.Release|x86.Deploy.0 = Release|x86
{A3B50604-0EAA-45F4-9991-39BD878E65A0}.Debug|Any CPU.ActiveCfg = Debug|x64
{A3B50604-0EAA-45F4-9991-39BD878E65A0}.Debug|Any CPU.Build.0 = Debug|x64
{A3B50604-0EAA-45F4-9991-39BD878E65A0}.Debug|ARM64.ActiveCfg = Debug|ARM64
{A3B50604-0EAA-45F4-9991-39BD878E65A0}.Debug|ARM64.Build.0 = Debug|ARM64
{A3B50604-0EAA-45F4-9991-39BD878E65A0}.Debug|x64.ActiveCfg = Debug|x64
{A3B50604-0EAA-45F4-9991-39BD878E65A0}.Debug|x64.Build.0 = Debug|x64
{A3B50604-0EAA-45F4-9991-39BD878E65A0}.Debug|x86.ActiveCfg = Debug|x86
{A3B50604-0EAA-45F4-9991-39BD878E65A0}.Debug|x86.Build.0 = Debug|x86
{A3B50604-0EAA-45F4-9991-39BD878E65A0}.Release|Any CPU.ActiveCfg = Release|x64
{A3B50604-0EAA-45F4-9991-39BD878E65A0}.Release|Any CPU.Build.0 = Release|x64
{A3B50604-0EAA-45F4-9991-39BD878E65A0}.Release|ARM64.ActiveCfg = Release|ARM64
{A3B50604-0EAA-45F4-9991-39BD878E65A0}.Release|ARM64.Build.0 = Release|ARM64
{A3B50604-0EAA-45F4-9991-39BD878E65A0}.Release|x64.ActiveCfg = Release|x64
{A3B50604-0EAA-45F4-9991-39BD878E65A0}.Release|x64.Build.0 = Release|x64
{A3B50604-0EAA-45F4-9991-39BD878E65A0}.Release|x86.ActiveCfg = Release|x86
{A3B50604-0EAA-45F4-9991-39BD878E65A0}.Release|x86.Build.0 = Release|x86
EndGlobalSection
GlobalSection(SolutionProperties) = preSolution
HideSolutionNode = FALSE
EndGlobalSection
GlobalSection(ExtensibilityGlobals) = postSolution
SolutionGuid = {CE9C717E-F1E7-45B8-AA13-092DCE0C9B28}
EndGlobalSection
EndGlobal

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/app.manifest
================================================

<?xml version="1.0" encoding="utf-8"?>
<assembly manifestVersion="1.0" xmlns="urn:schemas-microsoft-com:asm.v1">
  <assemblyIdentity version="1.0.0.0" name="ShiroProcessReporter.app"/>

  <compatibility xmlns="urn:schemas-microsoft-com:compatibility.v1">
    <application>
      <!-- The ID below informs the system that this application is compatible with OS features first introduced in Windows 10. 
      It is necessary to support features in unpackaged applications, for example the custom titlebar implementation.
      For more info see https://docs.microsoft.com/windows/apps/windows-app-sdk/use-windows-app-sdk-run-time#declare-os-compatibility-in-your-application-manifest -->
      <supportedOS Id="{8e0f7a12-bfb3-4fe8-b9a5-48fd50a15a9a}" />
    </application>
  </compatibility>
  
  <application xmlns="urn:schemas-microsoft-com:asm.v3">
    <windowsSettings>
      <dpiAwareness xmlns="http://schemas.microsoft.com/SMI/2016/WindowsSettings">PerMonitorV2</dpiAwareness>
    </windowsSettings>
  </application>
</assembly>

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/App.xaml
================================================

<?xml version="1.0" encoding="utf-8"?>

<Application
    x:Class="ShiroProcessReporter.App"
    xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
    xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"
    xmlns:local="using:ShiroProcessReporter"
    xmlns:helper="using:ShiroProcessReporter.Helper">
<Application.Resources>
<ResourceDictionary>
<ResourceDictionary.MergedDictionaries>
<XamlControlsResources xmlns="using:Microsoft.UI.Xaml.Controls" />

<!-- Other merged dictionaries here -->

</ResourceDictionary.MergedDictionaries>

<!-- Other app resources here -->

<helper:GlobalStateWrapper x:Key="globalStateWrapper"/>
</ResourceDictionary>
</Application.Resources>
</Application>

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/App.xaml.cs
================================================
﻿using H.NotifyIcon;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.UI.Xaml;
using ShiroProcessReporter.Services;
using ShiroProcessReporter.Components;
using ShiroProcessReporter.Helper;
using System;

namespace ShiroProcessReporter
{
public partial class App : Application
{
public static Window? MainWindow { get; private set; }

        public static IServiceProvider? ServiceProvider { get; private set; }

        public App()
        {
            this.InitializeComponent();

            AppLogHelper.ConfigureLogging();
        }

        private void ConfigureService(IServiceCollection services)
        {
            services.AddSingleton<MediaTraceService>();
            services.AddSingleton<ProcessTraceService>();
            services.AddSingleton<ReportService>();
            services.AddSingleton<TraceWorkerService>();
        }

        protected override void OnLaunched(LaunchActivatedEventArgs args)
        {
            Preferences.Initialize();
            var services = new ServiceCollection();
            ConfigureService(services);
            ServiceProvider = services.BuildServiceProvider();

            MainWindow = new MainView();

            MainWindow.Closed += (sender, args) =>
            {
                args.Handled = true;
                MainWindow.Hide();
            };

            MainWindow.Activate();
        }
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/MainView.xaml
================================================

<?xml version="1.0" encoding="utf-8"?>

<winuiex:WindowEx
x:Class="ShiroProcessReporter.MainView"
xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"
xmlns:local="using:ShiroProcessReporter"
xmlns:d="http://schemas.microsoft.com/expression/blend/2008"
xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006"
xmlns:components="using:ShiroProcessReporter.Components"
xmlns:layouts="using:ShiroProcessReporter.Layouts"
mc:Ignorable="d"
xmlns:winuiex="using:WinUIEx"
Title="Process Reporter Settings">
<Window.SystemBackdrop>
<MicaBackdrop />
</Window.SystemBackdrop>

    <Grid
        VerticalAlignment="Stretch"
        HorizontalAlignment="Stretch"
        Background="{ThemeResource ApplicationPageBackgroundThemeBrush}">
        <Grid.RowDefinitions>
            <RowDefinition Height="Auto" />
            <RowDefinition Height="*" />
        </Grid.RowDefinitions>
        <Grid
            x:Name="titleBar"
            Height="32"
            ColumnSpacing="16"
            Grid.Row="0">
            <Grid.ColumnDefinitions>
                <ColumnDefinition x:Name="LeftPaddingColumn" Width="0" />
                <ColumnDefinition x:Name="IconColumn" Width="Auto" />
                <ColumnDefinition x:Name="TitleColumn" Width="Auto" />
                <ColumnDefinition x:Name="RightPaddingColumn" Width="0" />
            </Grid.ColumnDefinitions>
            <Image
                Grid.Column="1"
                Width="16"
                Height="16"
                VerticalAlignment="Center"
                Source="../Assets/Hosts/Hosts.ico" />
            <TextBlock
                x:Name="AppTitleTextBlock"
                Grid.Column="2"
                VerticalAlignment="Center"
                Style="{StaticResource CaptionTextBlockStyle}" />
        </Grid>
        <layouts:Navigation Grid.Row="1" />
        <components:TrayIconView />
    </Grid>

</winuiex:WindowEx>

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/MainView.xaml.cs
================================================
using CommunityToolkit.Mvvm.ComponentModel;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.UI.Xaml;
using Microsoft.UI.Xaml.Controls;
using Microsoft.UI.Xaml.Controls.Primitives;
using Microsoft.UI.Xaml.Data;
using Microsoft.UI.Xaml.Input;
using Microsoft.UI.Xaml.Media;
using Microsoft.UI.Xaml.Navigation;
using ShiroProcessReporter.Services;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Runtime.InteropServices.WindowsRuntime;
using Windows.Foundation;
using Windows.Foundation.Collections;
using WinUIEx;

// To learn more about WinUI, the WinUI project structure,
// and more about our project templates, see: http://aka.ms/winui-project-info.

namespace ShiroProcessReporter
{
public sealed partial class MainView : WindowEx
{
private readonly TraceWorkerService \_traceWorkerService;

        public MainView()
        {
            this.InitializeComponent();
            _traceWorkerService = App.ServiceProvider.GetService<TraceWorkerService>();

            ExtendsContentIntoTitleBar = true;
            SetTitleBar(titleBar);
            var title = "Process Reporter Settings";
            Title = title;
            AppTitleTextBlock.Text = title;
            Activated += MainWindow_Activated;
        }

        private void MainWindow_Activated(object sender, WindowActivatedEventArgs args)
        {
            if (args.WindowActivationState == WindowActivationState.Deactivated)
            {
                AppTitleTextBlock.Foreground = (SolidColorBrush)App.Current.Resources["WindowCaptionForegroundDisabled"];
            }
            else
            {
                AppTitleTextBlock.Foreground = (SolidColorBrush)App.Current.Resources["WindowCaptionForeground"];
            }
        }

    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/NativeMethods.txt
================================================
DispatchMessage
GetClassName
GetForegroundWindow
GetMessage
GetProcessImageFileName
GetWindowText
GetWindowThreadProcessId
SetWinEventHook
TranslateMessage
EVENT_SYSTEM_FOREGROUND
WINEVENT_OUTOFCONTEXT
WINEVENT_SKIPOWNPROCESS

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/ShiroProcessReporter.csproj
================================================
﻿<Project Sdk="Microsoft.NET.Sdk">
<PropertyGroup>
<OutputType>WinExe</OutputType>
<TargetFramework>net8.0-windows10.0.19041.0</TargetFramework>
<TargetPlatformMinVersion>10.0.17763.0</TargetPlatformMinVersion>
<RootNamespace>ShiroProcessReporter</RootNamespace>
<ApplicationManifest>app.manifest</ApplicationManifest>
<Platforms>x86;x64;ARM64</Platforms>
<RuntimeIdentifiers>win-x86;win-x64;win-arm64</RuntimeIdentifiers>
<UseWinUI>true</UseWinUI>
<Nullable>enable</Nullable>
</PropertyGroup>
<ItemGroup>
<Content Include="Resources\TrayIcon\*.*" />
<Content Include="Resources\Images\*.*" />
</ItemGroup>
<ItemGroup>
<None Remove="Components\TrayIconView.xaml" />
<None Remove="Layouts\Navigation.xaml" />
<None Remove="NativeMethods.txt" />
<None Remove="Resources\Images\Bug.svg" />
<None Remove="Resources\Images\cloud-server-black.svg" />
<None Remove="Resources\Images\cloud-server-white.svg" />
<None Remove="Resources\Images\filter-black.svg" />
<None Remove="Resources\Images\filter-white.svg" />
<None Remove="Resources\Images\replace-black.svg" />
<None Remove="Resources\Images\replace-white.svg" />
<None Remove="Views\FilterRuleSettings.xaml" />
<None Remove="Views\LogViewer.xaml" />
</ItemGroup>
<ItemGroup>
<AdditionalFiles Include="NativeMethods.txt" />
</ItemGroup>

  <ItemGroup>
    <PackageReference Include="CommunityToolkit.Mvvm" Version="8.2.2" />
    <PackageReference Include="CommunityToolkit.WinUI" Version="7.1.2" />
    <PackageReference Include="CommunityToolkit.WinUI.Controls.SettingsControls" Version="8.0.240109" />
    <PackageReference Include="CommunityToolkit.WinUI.Converters" Version="8.0.240109" />
    <PackageReference Include="H.NotifyIcon.WinUI" Version="2.0.131" />
    <PackageReference Include="Microsoft.Extensions.DependencyInjection" Version="8.0.0" />
    <PackageReference Include="Microsoft.Extensions.Logging.Console" Version="8.0.0" />
    <PackageReference Include="Microsoft.Extensions.Logging.Debug" Version="8.0.0" />
    <PackageReference Include="Microsoft.Windows.CsWin32" Version="0.3.49-beta">
      <PrivateAssets>all</PrivateAssets>
      <IncludeAssets>runtime; build; native; contentfiles; analyzers; buildtransitive</IncludeAssets>
    </PackageReference>
    <PackageReference Include="Microsoft.WindowsAppSDK" Version="1.5.240428000" />
    <PackageReference Include="Microsoft.Windows.SDK.BuildTools" Version="10.0.22621.3233" />
    <PackageReference Include="WinUIEx" Version="2.3.4" />
    <Manifest Include="$(ApplicationManifest)" />
  </ItemGroup>

  <ItemGroup>
    <Page Update="Views\LogViewer.xaml">
      <Generator>MSBuild:Compile</Generator>
    </Page>
  </ItemGroup>

  <ItemGroup>
    <Page Update="Views\FilterRuleSettings.xaml">
      <Generator>MSBuild:Compile</Generator>
    </Page>
  </ItemGroup>

  <ItemGroup>
    <Page Update="Components\TrayIconView.xaml">
      <Generator>MSBuild:Compile</Generator>
    </Page>
  </ItemGroup>

  <ItemGroup>
    <Page Update="Layouts\Navigation.xaml">
      <Generator>MSBuild:Compile</Generator>
    </Page>
  </ItemGroup>
</Project>

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Components/TrayIconView.xaml
================================================

<?xml version="1.0" encoding="utf-8"?>

<UserControl
    x:Class="ShiroProcessReporter.Components.TrayIconView"
    xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
    xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"
    xmlns:local="using:ShiroProcessReporter.Components"
    xmlns:d="http://schemas.microsoft.com/expression/blend/2008"
    xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006"
    xmlns:tb="using:H.NotifyIcon"
    xmlns:converters="using:ShiroProcessReporter.Converters"
    xmlns:helper="using:ShiroProcessReporter.Helper"
    d:DataContext="{d:DesignInstance Type=local:TrayIconView}"
    mc:Ignorable="d">

    <UserControl.Resources>
        <converters:AppThemeToImageSourceConverter
            x:Key="AppThemeToImageSourceConverter"
            LightImage="ms-appx:///Resources/TrayIcon/flower.ico"
            DarkImage="ms-appx:///Resources/TrayIcon/flower_light.ico" />
    </UserControl.Resources>

    <tb:TaskbarIcon
        x:Name="TrayIcon"
        x:FieldModifier="public"
        ContextMenuMode="SecondWindow"
        IconSource="{x:Bind Path=helper:GlobalState.Instance.Theme, Converter={StaticResource AppThemeToImageSourceConverter}, Mode=OneWay}"
        LeftClickCommand="{x:Bind ShowHideWindowCommand}"
        NoLeftClickDelay="True"
        ToolTipText="ShiroProcessReporter">
        <tb:TaskbarIcon.ContextFlyout>
            <MenuFlyout>
                <MenuFlyoutItem Command="{x:Bind ShowHideWindowCommand}" Text="显示/隐藏设置" />
                <MenuFlyoutSeparator />
                <MenuFlyoutItem Command="{x:Bind ExitApplicationCommand}" Text="退出" />
            </MenuFlyout>
        </tb:TaskbarIcon.ContextFlyout>
    </tb:TaskbarIcon>

</UserControl>

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Components/TrayIconView.xaml.cs
================================================
using CommunityToolkit.Mvvm.ComponentModel;
using CommunityToolkit.Mvvm.Input;
using H.NotifyIcon;
using Microsoft.UI.Xaml;
using Microsoft.UI.Xaml.Controls;
using Microsoft.UI.Xaml.Controls.Primitives;
using Microsoft.UI.Xaml.Data;
using Microsoft.UI.Xaml.Input;
using Microsoft.UI.Xaml.Media;
using Microsoft.UI.Xaml.Media.Imaging;
using Microsoft.UI.Xaml.Navigation;
using ShiroProcessReporter.Helper;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Runtime.InteropServices.WindowsRuntime;
using Windows.Foundation;
using Windows.Foundation.Collections;

// To learn more about WinUI, the WinUI project structure,
// and more about our project templates, see: http://aka.ms/winui-project-info.

namespace ShiroProcessReporter.Components
{
[ObservableObject]
public sealed partial class TrayIconView : UserControl
{
[ObservableProperty]
private bool isWindowVisible;

        public TrayIconView()
        {
            this.InitializeComponent();

            this.ActualThemeChanged += OnThemeChanged;
        }

        private void OnThemeChanged(object sender, object? args)
        {
            GlobalState.Instance.Theme = Application.Current.RequestedTheme;
        }

        [RelayCommand]
        public void ShowHideWindow()
        {
            var window = App.MainWindow;
            if (window == null)
            {
                return;
            }

            if (window.Visible)
            {
                window?.Hide();
            }
            else
            {
                window?.Show();
            }
            IsWindowVisible = window?.Visible ?? false;
        }

        [RelayCommand]
        public void ExitApplication()
        {
            this.TrayIcon.Dispose();
            App.MainWindow?.Close();
        }
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Converters/AppThemeToImageSourceConverter.cs
================================================
﻿using Microsoft.UI.Xaml;
using Microsoft.UI.Xaml.Data;
using Microsoft.UI.Xaml.Media;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ShiroProcessReporter.Converters
{
public class AppThemeToImageSourceConverter : IValueConverter
{
public ImageSource LightImage { get; set; }
public ImageSource DarkImage { get; set; }

        public object Convert(object value, Type targetType, object parameter, string language)
        {
            if (ApplicationTheme.Dark == value as ApplicationTheme?)
            {
                return DarkImage;
            }

            return LightImage;
        }

        public object ConvertBack(object value, Type targetType, object parameter, string language)
        {
            throw new NotImplementedException();
        }
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Converters/LogLevelToColorConverter.cs
================================================
﻿using Microsoft.Extensions.Logging;
using Microsoft.UI;
using Microsoft.UI.Xaml;
using Microsoft.UI.Xaml.Data;
using Microsoft.UI.Xaml.Media;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Windows.UI;

namespace ShiroProcessReporter.Converters
{
public class LogLevelToColorConverter : IValueConverter
{
public object Convert(object value, Type targetType, object parameter, string language)
{
var brush = new SolidColorBrush(Colors.Gray); // 默认颜色
if (Application.Current.Resources.TryGetValue("SystemFillColorCriticalBrush", out var errorColor) && value is LogLevel logLevel)
{
switch (logLevel)
{
case LogLevel.Trace:
case LogLevel.Debug:
break;
case LogLevel.Information:
Application.Current.Resources.TryGetValue("SystemFillColorSuccessBrush", out var infoColor);
brush = (SolidColorBrush)infoColor;
break;
case LogLevel.Warning:
Application.Current.Resources.TryGetValue("SystemFillColorCautionBrush", out var warningColor);
brush = (SolidColorBrush)warningColor;
break;
case LogLevel.Error:
brush = (SolidColorBrush)errorColor;
break;
}
}
return brush;
}

        public object ConvertBack(object value, Type targetType, object parameter, string language)
        {
            throw new NotImplementedException();
        }
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Converters/ReplaceRulesDictionaryToArrayConverter.cs
================================================
﻿using Microsoft.UI.Xaml;
using Microsoft.UI.Xaml.Data;
using ShiroProcessReporter.Models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ShiroProcessReporter.Converters
{
public class ReplaceRulesDictionaryToArrayConverter : IValueConverter
{
public object Convert(object value, Type targetType, object parameter, string language)
{
if (value is not Dictionary<string, string>)
{
return new List<ReplaceRule>();
}

            Dictionary<string, string> pairs = value as Dictionary<string, string>;
            return pairs!.Select(item => new ReplaceRule { Original = item.Key, Replacement = item.Value }).ToArray();
        }

        public object ConvertBack(object value, Type targetType, object parameter, string language)
        {
            throw new NotImplementedException();
        }
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Extensions/MemoryLoggerExtension.cs
================================================
﻿using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Microsoft.Extensions.DependencyInjection.Extensions;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Logging.Configuration;
using ShiroProcessReporter.Providers;

namespace ShiroProcessReporter.Extensions
{
public static class MemoryLoggerExtension
{
public static ILoggingBuilder AddMemory(this ILoggingBuilder builder)
{
builder.AddConfiguration();

            builder.Services.TryAddEnumerable(
                ServiceDescriptor.Singleton<ILoggerProvider, MemoryLoggerProvider>());

            return builder;
        }
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Helper/AppLogHelper.cs
================================================
﻿using Microsoft.Extensions.Logging;
using ShiroProcessReporter.Extensions;
using ShiroProcessReporter.Models;
using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ShiroProcessReporter.Helper
{
public class AppLogHelper
{
public static ILoggerFactory Factory { get; set; }

        public static void ConfigureLogging()
        {
            Factory = LoggerFactory.Create(builder =>
            {
                builder
                    .AddDebug()
                    .AddConsole()
                    .AddMemory();
            });
        }
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Helper/DataContextWrapper.cs
================================================
﻿using ABI.System;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ShiroProcessReporter.Helper
{
public class DataContextWrapper<T>
{
public T Value { get; set; }

        public DataContextWrapper(T value)
        {
            this.Value = value;
        }
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Helper/Debouncer.cs
================================================
獵湩 ⁧ 祓瑳浥 ਻ 獵湩 ⁧ 祓瑳浥吮牨慥楤杮 ਻ 渊浡獥慰散匠楨潲牐捯獥剳灥牯整 ⹲ 效灬牥 ਻ 瀊扵楬 ⁣ 汣獡 ⁳ 敄潢湵散 ੲ੻††牰癩瑡 ⁥ 楔敭 ⁲ 瑟浩牥 ਻††牰癩瑡 ⁥ 楔敭灓湡张湩整癲污 ਻†† †瀠扵楬 ⁣ 敄潢湵散 ⡲ 楔敭灓湡椠瑮牥慶 ⥬ †笠  †††张湩整癲污㴠椠瑮牥慶㭬  †素 ਊ††異汢捩瘠楯 ⁤ 敄潢湵散䄨瑣潩 ⁮ 捡楴湯 ਩††੻††††⼯젠맧틻듑퓦뛚쪨욱ꏷ풬췲횣늹쾢믺쯙ૼ††††瑟浩牥⸿楄灳獯 ⡥ 㬩  †††ਠ††††⼯될붴튨뢻탶뗂뛄쪨욱૷††††瑟浩牥㴠渠睥吠浩牥弨㴠 ਾ††††੻††††††捡楴湯 ⤨਻††††††⼯혠킴췐뇪뫏ꏳ춬횣뚹쪨욱૷††††††瑟浩牥⸿楄灳獯 ⡥ 㬩  †††††张楴敭 ⁲‽畮汬 ਻††††ⱽ 渠汵 ⱬ 张湩整癲污 ‬ 楔敭灓湡䘮潲䵭汩楬敳潣摮 ⡳ ㄭ ⤩※⼯吠浩卥慰 ⹮ 牆浯楍汬獩捥湯獤 ⴨⤱넠쫭뚾쪨욱훷횻킴틐뒻૎††੽੽

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Helper/GlobalState.cs
================================================
﻿using CommunityToolkit.Mvvm.ComponentModel;
using Microsoft.UI.Xaml;
using ShiroProcessReporter.Models;
using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Microsoft.UI.Dispatching;

namespace ShiroProcessReporter.Helper
{
public partial class GlobalState : ObservableObject
{
[ObservableProperty]
private ApplicationTheme \_theme = Application.Current.RequestedTheme;

        [ObservableProperty]
        private ObservableCollection<LogEntry> _logs = [];

        public DispatcherQueue? LogViewDispatcherQueue { get; set; }

        private static GlobalState? _instance;
        public static GlobalState Instance => _instance ??= new GlobalState();
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Helper/GlobalStateWrapper.cs
================================================
﻿using ShiroProcessReporter.Models;
using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ShiroProcessReporter.Helper
{
public class GlobalStateWrapper
{
public ObservableCollection<LogEntry> Logs => GlobalState.Instance.Logs;
}
}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Helper/Preferences.cs
================================================
﻿using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Windows.ApplicationModel;
using Windows.Management.Core;

namespace ShiroProcessReporter.Helper
{
public static class Preferences
{
private static Windows.Storage.ApplicationDataContainer? LocalSettings;

        public static void Initialize()
        {
            LocalSettings = ApplicationDataManager.CreateForPackageFamily(Package.Current.Id.FamilyName).LocalSettings;
        }

        public static T? Get<T>(string key) where T : struct
        {
            if (LocalSettings is null)
            {
                Initialize();
            }

            if (string.IsNullOrWhiteSpace(key))
            {
                return null;
            }

            object value = LocalSettings.Values[key];

            if (value is null)
            {
                return null;
            }

            return (T)value;
        }

        public static T Get<T>(string key, T defaultValue)
        {
            if (LocalSettings is null)
            {
                Initialize();
            }

            if (string.IsNullOrWhiteSpace(key))
            {
                return defaultValue;
            }

            object value = LocalSettings.Values[key];

            if (value is null)
            {
                return defaultValue;
            }

            return (T)value;
        }

        public static void Set<T>(string key, T value)
        {
            if (LocalSettings is null)
            {
                Initialize();
            }

            if (string.IsNullOrWhiteSpace(key))
            {
                return;
            }

            LocalSettings.Values[key] = value;
        }

    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Layouts/Navigation.xaml
================================================

<?xml version="1.0" encoding="utf-8"?>

<Page
    x:Class="ShiroProcessReporter.Layouts.Navigation"
    xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
    xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"
    xmlns:local="using:ShiroProcessReporter.Layouts"
    xmlns:d="http://schemas.microsoft.com/expression/blend/2008"
    xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006"
    xmlns:converters="using:ShiroProcessReporter.Converters"
    xmlns:helper="using:ShiroProcessReporter.Helper"
    mc:Ignorable="d">

    <Page.Resources>
        <converters:AppThemeToImageSourceConverter
            x:Key="ServerIconConverter"
            LightImage="ms-appx:///Resources/Images/cloud-server-black.svg"
            DarkImage="ms-appx:///Resources/Images/cloud-server-white.svg" />
        <converters:AppThemeToImageSourceConverter
            x:Key="ReplaceIconConverter"
            LightImage="ms-appx:///Resources/Images/replace-black.svg"
            DarkImage="ms-appx:///Resources/Images/replace-white.svg" />
        <converters:AppThemeToImageSourceConverter
            x:Key="FilterIconConverter"
            LightImage="ms-appx:///Resources/Images/filter-black.svg"
            DarkImage="ms-appx:///Resources/Images/filter-white.svg" />
        <converters:AppThemeToImageSourceConverter
            x:Key="LogsIconConverter"
            LightImage="ms-appx:///Resources/Images/bug-black.svg"
            DarkImage="ms-appx:///Resources/Images/bug-white.svg" />
    </Page.Resources>

    <NavigationView
        x:Name="NavView"
        Background="Transparent"
        IsSettingsVisible="False"
        IsBackEnabled="False"
        IsBackButtonVisible="Collapsed"
        IsPaneToggleButtonVisible="False"
        Loaded="NavigationView_Loaded"
        ItemInvoked="NavigationView_ItemInvoked">
        <NavigationView.PaneHeader>
            <StackPanel VerticalAlignment="Center" HorizontalAlignment="Center">
                <Image Source="ms-appx:///Resources/Images/mix_space.png" Width="300" Height="200" />
            </StackPanel>
        </NavigationView.PaneHeader>
        <NavigationView.Header>
            <TextBlock Style="{StaticResource TitleTextBlockStyle}" Text="{x:Bind Title, Mode=OneWay}" />
        </NavigationView.Header>
        <NavigationView.MenuItems>
            <NavigationViewItem Tag="ShiroProcessReporter.Views.EndpointSettings" Content="Mix Space Endpoint">
                <NavigationViewItem.Icon>
                    <ImageIcon Source="{x:Bind helper:GlobalState.Instance.Theme, Converter={StaticResource ServerIconConverter}, Mode=OneWay}">
                        <ImageIcon.RenderTransform>
                            <ScaleTransform ScaleX="1.25" ScaleY="1.25" />
                        </ImageIcon.RenderTransform>
                    </ImageIcon>
                </NavigationViewItem.Icon>
            </NavigationViewItem>
            <NavigationViewItem Tag="ShiroProcessReporter.Views.ReplaceRuleSettings" Content="Replace Rules">
                <NavigationViewItem.Icon>
                    <ImageIcon Source="{x:Bind helper:GlobalState.Instance.Theme, Converter={StaticResource ReplaceIconConverter}, Mode=OneWay}" >
                        <ImageIcon.RenderTransform>
                            <ScaleTransform ScaleX="1.1" ScaleY="1.1" />
                        </ImageIcon.RenderTransform>
                    </ImageIcon>
                </NavigationViewItem.Icon>
            </NavigationViewItem>
            <NavigationViewItem Tag="ShiroProcessReporter.Views.FilterRuleSettings" Content="Filter Rules">
                <NavigationViewItem.Icon>
                    <ImageIcon Source="{x:Bind helper:GlobalState.Instance.Theme, Converter={StaticResource FilterIconConverter}, Mode=OneWay}">
                        <ImageIcon.RenderTransform>
                            <ScaleTransform ScaleX="1.1" ScaleY="1.1" />
                        </ImageIcon.RenderTransform>
                    </ImageIcon>
                </NavigationViewItem.Icon>
            </NavigationViewItem>
            <NavigationViewItem Tag="ShiroProcessReporter.Views.LogViewer" Content="Logs">
                <NavigationViewItem.Icon>
                    <ImageIcon Source="{x:Bind helper:GlobalState.Instance.Theme, Converter={StaticResource LogsIconConverter}, Mode=OneWay}">
                        <ImageIcon.RenderTransform>
                            <ScaleTransform ScaleX="1.25" ScaleY="1.25" />
                        </ImageIcon.RenderTransform>
                    </ImageIcon>
                </NavigationViewItem.Icon>
            </NavigationViewItem>
        </NavigationView.MenuItems>
        <NavigationView.Content>
            <Frame x:Name="RouterView" />
        </NavigationView.Content>
    </NavigationView>

</Page>

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Layouts/Navigation.xaml.cs
================================================
using CommunityToolkit.Mvvm.ComponentModel;
using Microsoft.UI.Xaml;
using Microsoft.UI.Xaml.Controls;
using Microsoft.UI.Xaml.Controls.Primitives;
using Microsoft.UI.Xaml.Data;
using Microsoft.UI.Xaml.Input;
using Microsoft.UI.Xaml.Media;
using Microsoft.UI.Xaml.Media.Animation;
using Microsoft.UI.Xaml.Navigation;
using ShiroProcessReporter.Views;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Runtime.InteropServices.WindowsRuntime;
using Windows.Foundation;
using Windows.Foundation.Collections;
using Windows.UI.ApplicationSettings;

namespace ShiroProcessReporter.Layouts
{
[ObservableObject]
public sealed partial class Navigation : Page
{
public Navigation()
{
this.InitializeComponent();
}

        [ObservableProperty]
        private string _title = "";

        private void NavigationView_Loaded(object sender, RoutedEventArgs e)
        {
            NavView.SelectedItem = NavView.MenuItems[0];
            Navigate(typeof(EndpointSettings), new EntranceNavigationTransitionInfo());
        }

        private void NavigationView_ItemInvoked(NavigationView sender, NavigationViewItemInvokedEventArgs args)
        {
            if (args.IsSettingsInvoked == true)
            {
                //Navigate(typeof(SettingsPage), args.RecommendedNavigationTransitionInfo);
            }
            else if (args.InvokedItemContainer != null)
            {
                Type? navPageType = Type.GetType(args.InvokedItemContainer.Tag.ToString());
                if (navPageType is null)
                {
                    return;
                }
                Navigate(navPageType, args.RecommendedNavigationTransitionInfo);
            }
        }

        private void Navigate(Type navPageType, NavigationTransitionInfo transitionInfo)
        {
            Type preNavPageType = RouterView.CurrentSourcePageType;
            if (navPageType is not null && !Type.Equals(preNavPageType, navPageType))
            {
                if (RouterView.Content is Page currentPage && currentPage is IDisposable disposable)
                {
                    disposable.Dispose();
                }
                RouterView.Navigate(navPageType, null, transitionInfo);
                Title = (NavView.SelectedItem as NavigationViewItem).Content.ToString();
            }
        }
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Loggers/MemoryLogger.cs
================================================
﻿using Microsoft.Extensions.Logging;
using Microsoft.UI.Dispatching;
using ShiroProcessReporter.Helper;
using ShiroProcessReporter.Models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ShiroProcessReporter.Loggers
{
public sealed class MemoryLogger : ILogger
{
public IDisposable? BeginScope<TState>(TState state) where TState : notnull => default!;

        public bool IsEnabled(LogLevel logLevel) => true;

        public void Log<TState>(
            LogLevel logLevel,
            EventId eventId,
            TState state,
            Exception? exception,
            Func<TState, Exception?, string> formatter)
        {
            var dispatcher = GlobalState.Instance.LogViewDispatcherQueue;

            var entry = new LogEntry
            {
                LogLevel = logLevel,
                Message = $"[{DateTime.Now} {logLevel}]  {formatter(state, exception)}",
            };

            if (dispatcher == null)
            {
                AddLog(entry);
            }
            else
            {
                dispatcher.TryEnqueue(() =>
                {
                    AddLog(entry);
                });
            }
        }

        private void AddLog(LogEntry entry)
        {
            GlobalState.Instance.Logs.Add(entry);

            while (GlobalState.Instance.Logs.Count > 1000)
            {
                GlobalState.Instance.Logs.RemoveAt(0);
            }
        }
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Models/FilterRule.cs
================================================
﻿using CommunityToolkit.Mvvm.ComponentModel;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;

namespace ShiroProcessReporter.Models
{
public partial class FilterRule : ObservableObject
{
[ObservableProperty]
[NotifyPropertyChangedFor(nameof(FilterRuleVaild))]
[NotifyPropertyChangedFor(nameof(IsOriginalVaild))]
public string \_original;

        public FilterRule Clone()
        {
            return new FilterRule
            {
                Original = Original,
            };
        }

        public bool FilterRuleVaild => IsOriginalVaild;

        public bool IsOriginalVaild => ValidateOriginal();

        private bool ValidateOriginal()
        {
            if (string.IsNullOrWhiteSpace(Original))
            {
                return false;
            }

            try
            {
                var regex = new Regex(Original);
                return true;
            }
            catch (ArgumentException)
            {
                return false;
            }
        }
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Models/LogEntry.cs
================================================
﻿using CommunityToolkit.Mvvm.ComponentModel;
using Microsoft.Extensions.Logging;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;

namespace ShiroProcessReporter.Models
{
public partial class LogEntry
{
public string Message { get; set; }
public LogLevel LogLevel { get; set; }
}
}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Models/ReplaceRule.cs
================================================
﻿using CommunityToolkit.Mvvm.ComponentModel;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;

namespace ShiroProcessReporter.Models
{
public partial class ReplaceRule : ObservableObject
{
[ObservableProperty]
[NotifyPropertyChangedFor(nameof(ReplaceRuleVaild))]
[NotifyPropertyChangedFor(nameof(IsOriginalVaild))]
public string \_original;

        [ObservableProperty]
        [NotifyPropertyChangedFor(nameof(ReplaceRuleVaild))]
        [NotifyPropertyChangedFor(nameof(IsOriginalVaild))]
        public string _replacement;

        public ReplaceRule Clone()
        {
            return new ReplaceRule
            {
                Original = Original,
                Replacement = Replacement,
            };
        }

        public bool ReplaceRuleVaild => IsOriginalVaild && !string.IsNullOrEmpty(Replacement);

        public bool IsOriginalVaild => ValidateOriginal();

        private bool ValidateOriginal()
        {
            if (string.IsNullOrWhiteSpace(Original))
            {
                return false;
            }

            try
            {
                var regex = new Regex(Original);
                return true;
            }
            catch (ArgumentException)
            {
                return false;
            }
        }
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Properties/PublishProfiles/win-arm64.pubxml
================================================
﻿<?xml version="1.0" encoding="utf-8"?>

<!--
https://go.microsoft.com/fwlink/?LinkID=208121.
-->
<Project ToolsVersion="4.0" xmlns="http://schemas.microsoft.com/developer/msbuild/2003">
  <PropertyGroup>
    <PublishProtocol>FileSystem</PublishProtocol>
    <Platform>ARM64</Platform>
    <RuntimeIdentifier>win-arm64</RuntimeIdentifier>
    <PublishDir>bin\$(Configuration)\$(TargetFramework)\$(RuntimeIdentifier)\publish\</PublishDir>
    <SelfContained>true</SelfContained>
    <PublishSingleFile>False</PublishSingleFile>
    <PublishReadyToRun Condition="'$(Configuration)' == 'Debug'">False</PublishReadyToRun>
    <PublishReadyToRun Condition="'$(Configuration)' != 'Debug'">True</PublishReadyToRun>
   <!-- 
    See https://github.com/microsoft/CsWinRT/issues/373
    <PublishTrimmed>True</PublishTrimmed>
    -->
  </PropertyGroup>
</Project>

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Properties/PublishProfiles/win-x64.pubxml
================================================
﻿<?xml version="1.0" encoding="utf-8"?>

<!--
https://go.microsoft.com/fwlink/?LinkID=208121.
-->
<Project ToolsVersion="4.0" xmlns="http://schemas.microsoft.com/developer/msbuild/2003">
  <PropertyGroup>
    <PublishProtocol>FileSystem</PublishProtocol>
    <Platform>x64</Platform>
    <RuntimeIdentifier>win-x64</RuntimeIdentifier>
    <PublishDir>bin\$(Configuration)\$(TargetFramework)\$(RuntimeIdentifier)\publish\</PublishDir>
    <SelfContained>true</SelfContained>
    <PublishSingleFile>False</PublishSingleFile>
    <PublishReadyToRun Condition="'$(Configuration)' == 'Debug'">False</PublishReadyToRun>
    <PublishReadyToRun Condition="'$(Configuration)' != 'Debug'">True</PublishReadyToRun>
   <!-- 
    See https://github.com/microsoft/CsWinRT/issues/373
    <PublishTrimmed>True</PublishTrimmed>
    -->
  </PropertyGroup>
</Project>

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Properties/PublishProfiles/win-x86.pubxml
================================================
﻿<?xml version="1.0" encoding="utf-8"?>

<!--
https://go.microsoft.com/fwlink/?LinkID=208121.
-->
<Project ToolsVersion="4.0" xmlns="http://schemas.microsoft.com/developer/msbuild/2003">
  <PropertyGroup>
    <PublishProtocol>FileSystem</PublishProtocol>
    <Platform>x86</Platform>
    <RuntimeIdentifier>win-x86</RuntimeIdentifier>
    <PublishDir>bin\$(Configuration)\$(TargetFramework)\$(RuntimeIdentifier)\publish\</PublishDir>
    <SelfContained>true</SelfContained>
    <PublishSingleFile>False</PublishSingleFile>
    <PublishReadyToRun Condition="'$(Configuration)' == 'Debug'">False</PublishReadyToRun>
    <PublishReadyToRun Condition="'$(Configuration)' != 'Debug'">True</PublishReadyToRun>
   <!-- 
    See https://github.com/microsoft/CsWinRT/issues/373
    <PublishTrimmed>True</PublishTrimmed>
    -->
  </PropertyGroup>
</Project>

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Providers/MemoryLoggerProvider.cs
================================================
﻿using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Options;
using ShiroProcessReporter.Loggers;
using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ShiroProcessReporter.Providers
{
public sealed class MemoryLoggerProvider : ILoggerProvider
{
private readonly ConcurrentDictionary<string, MemoryLogger> \_loggers =
new(StringComparer.OrdinalIgnoreCase);

        public ILogger CreateLogger(string categoryName) =>
            _loggers.GetOrAdd(categoryName, name => new MemoryLogger());

        public void Dispose()
        {
            _loggers.Clear();
        }
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Services/MediaTraceService.cs
================================================
﻿using System;
using Windows.Media.Control;
using Microsoft.Extensions.Logging;
using ShiroProcessReporter.Helper;

namespace ShiroProcessReporter.Services;

public class MediaTraceService
{
public delegate void MediaPlaybackChangeHandler(GlobalSystemMediaTransportControlsSessionMediaProperties properties,
GlobalSystemMediaTransportControlsSessionPlaybackStatus status, string processName);

    private readonly ILogger<MediaTraceService> _logger;

    private GlobalSystemMediaTransportControlsSessionManager? _manager;

    private string _processName = string.Empty;

    private GlobalSystemMediaTransportControlsSessionMediaProperties? _properties;

    private GlobalSystemMediaTransportControlsSessionPlaybackStatus _status;

    public MediaTraceService()
    {
        _logger = AppLogHelper.Factory.CreateLogger<MediaTraceService>();
        _properties = null;
        _status = GlobalSystemMediaTransportControlsSessionPlaybackStatus.Stopped;
        InitManager();
    }

    public event MediaPlaybackChangeHandler OnMediaPlaybackChanged;

    private async void InitManager()
    {
        _manager = await GlobalSystemMediaTransportControlsSessionManager.RequestAsync();
        if (_manager is null)
        {
            _logger.LogError("Failed to get MediaSessionManager");
            return;
        }

        // 初始化触发一次
        MediaSessionChanged(_manager, null);
        _manager.CurrentSessionChanged += MediaSessionChanged;
    }

    private void MediaSessionChanged(GlobalSystemMediaTransportControlsSessionManager manager,
        CurrentSessionChangedEventArgs? args)
    {
        var session = manager.GetCurrentSession();
        if (session is null)
        {
            _processName = string.Empty;
            _properties = null;
            _status = GlobalSystemMediaTransportControlsSessionPlaybackStatus.Stopped;

            _logger.LogWarning("Media session is null");
            return;
        }

        // 初始化触发一次
        MediaPropertiesChanged(session, null);
        PlaybackInfoChanged(session, null);

        _processName = session.SourceAppUserModelId;
        session.MediaPropertiesChanged += MediaPropertiesChanged;
        session.PlaybackInfoChanged += PlaybackInfoChanged;
    }

    private void PlaybackInfoChanged(GlobalSystemMediaTransportControlsSession sender,
        PlaybackInfoChangedEventArgs? args)
    {
        var playbackInfo = sender.GetPlaybackInfo();
        _status = playbackInfo.PlaybackStatus;

        _logger.LogInformation("Play status change to {status}", playbackInfo.PlaybackStatus);

        CallEvent();
    }

    private async void MediaPropertiesChanged(GlobalSystemMediaTransportControlsSession sender,
        MediaPropertiesChangedEventArgs? args)
    {
        try
        {
            var mediaProperities = await sender.TryGetMediaPropertiesAsync();
            if (mediaProperities is null)
            {
                _logger.LogError("Media properties is null");
                return;
            }

            _properties = mediaProperities;

            _logger.LogInformation("Playing {artist} - {title}", mediaProperities.Artist, mediaProperities.Title);

            CallEvent();
        }
        catch (Exception)
        {
            _logger.LogWarning("Failed to get media properties");
        }
    }

    private void CallEvent()
    {
        if (_properties is not null)
            if (OnMediaPlaybackChanged is not null)
                OnMediaPlaybackChanged(_properties, _status, _processName);
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Services/ProcessTraceService.cs
================================================
﻿using System.Buffers;
using System.Diagnostics;
using Windows.Win32.Foundation;
using Windows.Win32.UI.Accessibility;
using Microsoft.Extensions.Logging;
using static Windows.Win32.PInvoke;
using ShiroProcessReporter;
using Microsoft.Extensions.DependencyInjection;
using System;
using ShiroProcessReporter.Helper;

namespace ShiroProcessReporter.Services;

// Reference: https://github.com/walterlv/Walterlv.ForegroundWindowMonitor

public class ProcessTraceService
{
public delegate void FrontWindowChangeHandler(string windowTitle, string processName);

    private readonly ILogger<ProcessTraceService> _logger;
    private readonly HWINEVENTHOOK _hookHandle;

    public ProcessTraceService()
    {
        _logger = AppLogHelper.Factory.CreateLogger<ProcessTraceService>();

        _hookHandle = SetWinEventHook(
            EVENT_SYSTEM_FOREGROUND,
            EVENT_SYSTEM_FOREGROUND,
            HMODULE.Null,
            OnFrontWindowChange,
            0, 0,
            WINEVENT_OUTOFCONTEXT | WINEVENT_SKIPOWNPROCESS);

        // 开启消息循环，以便 WinEventProc 能够被调用。
        if (GetMessage(out var lpMsg, default, default, default))
        {
            TranslateMessage(in lpMsg);
            DispatchMessage(in lpMsg);
        }
    }

    public event FrontWindowChangeHandler OnFrontWindowChanged;

    ~ProcessTraceService()
    {
        UnhookWinEvent(_hookHandle);
    }

    private static void OnFrontWindowChange(HWINEVENTHOOK eventHook, uint @event, HWND hwnd, int idObject, int idChild,
        uint idEventThread, uint dwmsEventTime)
    {
        var currentWindow = GetForegroundWindow();

        var processId = GetProcessIdCore(currentWindow);

        var processName = Process.GetProcessById((int)processId).ProcessName;

        var windowTitle = CallWin32ToGetPWSTR(512, (p, l) => GetWindowText(currentWindow, p, l));

        var service = App.ServiceProvider.GetService<ProcessTraceService>();

        if (service?.OnFrontWindowChanged is not null) service.OnFrontWindowChanged(windowTitle, processName);
    }

    private static unsafe uint GetProcessIdCore(HWND hWnd)
    {
        uint pid = 0;
        GetWindowThreadProcessId(hWnd, &pid);
        return pid;
    }

    private static unsafe string CallWin32ToGetPWSTR(int bufferLength, Func<PWSTR, int, int> getter)
    {
        var buffer = ArrayPool<char>.Shared.Rent(bufferLength);
        try
        {
            fixed (char* ptr = buffer)
            {
                getter(ptr, bufferLength);
                return new string(buffer, 0, Array.IndexOf(buffer, '\0'));
            }
        }
        finally
        {
            ArrayPool<char>.Shared.Return(buffer);
        }
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Services/ReportService.cs
================================================
﻿using System.Net.Http.Json;
using System.Text.Json;
using System.Text.RegularExpressions;
using Windows.Media.Control;
using Microsoft.Extensions.Logging;
using System;
using System.Net.Http;
using System.Collections.Generic;
using Windows.ApplicationModel;
using ShiroProcessReporter.Helper;
using System.Linq;
using ShiroProcessReporter.Models;
using Windows.ApplicationModel.Appointments.AppointmentsProvider;

namespace ShiroProcessReporter.Services;

public class ReportService
{
private readonly HttpClient \_httpClient = new()
{
Timeout = TimeSpan.FromSeconds(5)
};

    private readonly ILogger<ReportService> _logger;
    private readonly string ApiKeyDefault = string.Empty;

    private readonly List<ReplaceRule> BuiltinReplaceRules = new()
    {
        new ReplaceRule{ Original = @"\.[Ee][Xx][Ee]", Replacement = "" }, // 删除exe后缀

        new ReplaceRule{ Original = "[Ee]xplorer", Replacement = "explorer" },

        new ReplaceRule{ Original = "msedge", Replacement = "Microsoft Edge" },
        new ReplaceRule{ Original = "WINWORD", Replacement = "Microsoft Word" },
        new ReplaceRule{ Original = "EXCEL", Replacement = "Microsoft Excel" },
        new ReplaceRule{ Original = "POWERPNT", Replacement = "Microsoft PowerPoint" },
        new ReplaceRule{ Original = "ONENOTE", Replacement = "Microsoft OneNote" },

        new ReplaceRule{ Original = "idea64", Replacement = "IntelliJ IDEA" },
        new ReplaceRule{ Original = "goland64", Replacement = "GoLand" },
        new ReplaceRule{ Original = "pycharm64", Replacement = "PyCharm" },

        new ReplaceRule{ Original = "GitHubDesktop", Replacement = "GitHub Desktop" },
        new ReplaceRule{ Original = "chrome", Replacement = "Chrome" },
    };

    private readonly string EndpointDefault = string.Empty;
    private readonly List<FilterRule> FilterRulesDefault = [];
    private readonly List<ReplaceRule> ReplaceRulesDefault = [];

    private readonly string IdApiKey = "api_key";

    private readonly string IdEndpoint = "endpoint";

    private readonly string IdFilterRules = "filter_rules";

    private readonly string IdReplaceRules = "replace_rules";

    private readonly string UserAgent
        = $"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 ProcessReporter/{Package.Current.Id.Version}";

    public ReportService()
    {
        _logger = AppLogHelper.Factory.CreateLogger<ReportService>();
        _httpClient.DefaultRequestHeaders.Add("User-Agent", UserAgent);
    }

    public string Endpoint
    {
        get => Preferences.Get(IdEndpoint, EndpointDefault);
        set => Preferences.Set(IdEndpoint, value);
    }

    public string ApiKey
    {
        get => Preferences.Get(IdApiKey, ApiKeyDefault);
        set => Preferences.Set(IdApiKey, value);
    }

    public List<ReplaceRule> ReplaceRules
    {
        get => JsonSerializer.Deserialize<List<ReplaceRule>>(Preferences.Get(IdReplaceRules, "[]")) ?? ReplaceRulesDefault;
        set => Preferences.Set(IdReplaceRules, JsonSerializer.Serialize(value));
    }

    public List<FilterRule> FilterRules
    {
        get => JsonSerializer.Deserialize<List<FilterRule>>(Preferences.Get(IdFilterRules, "[]")) ?? FilterRulesDefault;
        set => Preferences.Set(IdFilterRules, JsonSerializer.Serialize(value));
    }

    private List<ReplaceRule> MergedReplaceRules => [.. BuiltinReplaceRules, .. ReplaceRules];

    public async void ReportProcess(string windowTitle, string processName)
    {
        if (string.IsNullOrEmpty(Endpoint))
        {
            _logger.LogInformation("Endpoint is empty, skip report");
            return;
        }

        if (string.IsNullOrEmpty(ApiKey))
        {
            _logger.LogInformation("Api key is empty, skip report");
            return;
        }

        var timeStamp = DateTimeOffset.Now.ToUnixTimeMilliseconds();

        foreach (var rule in FilterRules)
        {
            var regex = new Regex(Regex.Unescape(rule.Original));
            if (regex.IsMatch(processName))
            {
                _logger.LogInformation("Process {process} matched filter rule {rule}", processName, rule.Original);
                return;
            }
        }

        foreach (var rule in MergedReplaceRules)
        {
            var regex = new Regex(Regex.Unescape(rule.Original));
            processName = regex.Replace(processName, rule.Replacement);
        }

        try
        {
            var response = await _httpClient.PostAsJsonAsync(Endpoint, new Dictionary<string, object>
            {
                { "timestamp", timeStamp },
                { "process", processName },
                { "title", windowTitle },
                { "key", ApiKey }
            });

            response.EnsureSuccessStatusCode();

            var dataStr = await response.Content.ReadAsStringAsync();

            response.Dispose();

            if (string.IsNullOrEmpty(dataStr)) throw new Exception("Api server return empty content");

            var data = JsonSerializer.Deserialize<Dictionary<string, object>>(dataStr) ??
                       throw new Exception("Invaild response");

            if (data.TryGetValue("ok", out var eOk))
            {
                var ok = ((JsonElement?)eOk).Value.GetInt32();
                if (ok == 0)
                    if (data.TryGetValue("message", out var eMsg))
                    {
                        var msg = ((JsonElement?)eMsg).Value.GetString();
                        throw new Exception(msg);
                    }
            }

            _logger.LogInformation("Report process success");
        }
        catch (Exception ex)
        {
            _logger.LogError("Report failed, reason: {message}", ex.Message);
        }
    }

    public async void ReportMedia(GlobalSystemMediaTransportControlsSessionMediaProperties properties,
        GlobalSystemMediaTransportControlsSessionPlaybackStatus status, string processName)
    {
        if (status != GlobalSystemMediaTransportControlsSessionPlaybackStatus.Playing) return;

        if (string.IsNullOrEmpty(Endpoint))
        {
            _logger.LogInformation("Endpoint is empty, skip report");
            return;
        }

        if (string.IsNullOrEmpty(ApiKey))
        {
            _logger.LogInformation("Api key is empty, skip report");
            return;
        }

        if (string.IsNullOrEmpty(properties.Artist))
        {
            _logger.LogWarning("No artist info found, don't think it's a music program");
            return;
        }

        if (string.IsNullOrEmpty(properties.Title))
        {
            _logger.LogWarning("No title info found, don't think it's a music program");
            return;
        }

        var timeStamp = DateTimeOffset.Now.ToUnixTimeMilliseconds();

        foreach (var rule in FilterRules)
        {
            var regex = new Regex(Regex.Unescape(rule.Original));
            if (regex.IsMatch(processName))
            {
                _logger.LogInformation("Process {process} matched filter rule {rule}", processName, rule.Original);
                return;
            }
        }

        foreach (var rule in MergedReplaceRules)
        {
            var regex = new Regex(Regex.Unescape(rule.Original));
            processName = regex.Replace(processName, rule.Replacement);
        }

        try
        {
            var response = await _httpClient.PostAsJsonAsync(Endpoint, new Dictionary<string, object>
            {
                { "timestamp", timeStamp },
                { "process", processName },
                {
                    "media", new Dictionary<string, string>
                    {
                        { "title", properties.Title },
                        { "artist", properties.Artist }
                    }
                },
                { "key", ApiKey }
            });

            response.EnsureSuccessStatusCode();

            var dataStr = await response.Content.ReadAsStringAsync();

            response.Dispose();

            if (string.IsNullOrEmpty(dataStr)) throw new Exception("Api server return empty content");

            var data = JsonSerializer.Deserialize<Dictionary<string, object>>(dataStr) ??
                       throw new Exception("Invaild response");

            if (data.TryGetValue("ok", out var eOk))
            {
                var ok = ((JsonElement?)eOk).Value.GetInt32();
                if (ok == 0)
                    if (data.TryGetValue("message", out var eMsg))
                    {
                        var msg = ((JsonElement?)eMsg).Value.GetString();
                        throw new Exception(msg);
                    }
            }

            _logger.LogInformation("Report media success");
        }
        catch (Exception ex)
        {
            _logger.LogError("Report failed, reason: {message}", ex.Message);
        }
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Services/TraceWorkerService.cs
================================================
﻿using Windows.Media.Control;
using Microsoft.Extensions.Logging;
using ShiroProcessReporter.Helper;
using ShiroProcessReporter;
using Microsoft.Extensions.DependencyInjection;

namespace ShiroProcessReporter.Services;

public class TraceWorkerService
{
private readonly ILogger<TraceWorkerService> \_logger;

    private readonly MediaTraceService _mediaTraceService;

    private readonly ProcessTraceService _processTraceService;

    private readonly ReportService _reportService;

    public TraceWorkerService()
    {
        _logger = AppLogHelper.Factory.CreateLogger<TraceWorkerService>();
        _mediaTraceService = App.ServiceProvider.GetService<MediaTraceService>();
        _processTraceService = App.ServiceProvider.GetService<ProcessTraceService>();
        _reportService = App.ServiceProvider.GetService<ReportService>();

        _mediaTraceService.OnMediaPlaybackChanged += OnMediaPlaybackChanged;
        _processTraceService.OnFrontWindowChanged += OnFrontWindowChanged;
    }

    private void OnFrontWindowChanged(string windowTitle, string processName)
    {
        _logger.LogInformation("Working on {title} - {process}", windowTitle, processName);
        _reportService.ReportProcess(windowTitle, processName);
    }

    private void OnMediaPlaybackChanged(GlobalSystemMediaTransportControlsSessionMediaProperties properties,
        GlobalSystemMediaTransportControlsSessionPlaybackStatus status, string processName)
    {
        _logger.LogInformation("Now Playing: {Artist} - {Title} => {status}", properties.Artist, properties.Title,
            status);
        _reportService.ReportMedia(properties, status, processName);
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Views/EndpointSettings.xaml
================================================

<?xml version="1.0" encoding="utf-8"?>

<Page
    x:Class="ShiroProcessReporter.Views.EndpointSettings"
    xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
    xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"
    xmlns:local="using:ShiroProcessReporter.Views"
    xmlns:controls="using:CommunityToolkit.WinUI.Controls"
    xmlns:helpers="using:ShiroProcessReporter.Helper"
    xmlns:d="http://schemas.microsoft.com/expression/blend/2008"
    xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006"
    mc:Ignorable="d"
    Background="Transparent">

    <StackPanel HorizontalAlignment="Stretch" VerticalAlignment="Top" Padding="20" Spacing="4">
        <ContentDialog
            x:Name="EndpointEditDialog"
            PrimaryButtonText="Save"
            SecondaryButtonText="Cancel"
            DefaultButton="Primary">
            <ContentDialog.DataContext />
            <TextBox x:Name="DialogInput" Text="{Binding Value, Mode=TwoWay, UpdateSourceTrigger=PropertyChanged}" />
        </ContentDialog>

        <controls:SettingsCard
            Header="Mix Space Server URL"
            Description="Your Mix Space URL, eg. https://mxspace.mydomain.com/api/v2/fn/ps/update">
            <controls:SettingsCard.HeaderIcon>
                <SymbolIcon Symbol="Link" />
            </controls:SettingsCard.HeaderIcon>
            <Button Content="Edit" Style="{StaticResource AccentButtonStyle}" Click="EditEndpoint" />
        </controls:SettingsCard>
        <controls:SettingsCard
            Header="Mix Space Server API Key"
            Description="Your serverless key">
            <controls:SettingsCard.HeaderIcon>
                <SymbolIcon Symbol="Permissions" />
            </controls:SettingsCard.HeaderIcon>
            <Button Content="Edit" Style="{StaticResource AccentButtonStyle}" Click="EditApiKey" />
        </controls:SettingsCard>
        <controls:SettingsCard
            Header="Have any problem?">
            <controls:SettingsCard.HeaderIcon>
                <SymbolIcon Symbol="Help" />
            </controls:SettingsCard.HeaderIcon>
            <controls:SettingsCard.Description>
                <HyperlinkButton NavigateUri="https://mx-space.js.org/themes/shiro/extra#%E6%88%91%E7%9A%84%E5%8A%A8%E6%80%81" Content="Get more info from document" />
            </controls:SettingsCard.Description>
        </controls:SettingsCard>
    </StackPanel>

</Page>

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Views/EndpointSettings.xaml.cs
================================================
using Microsoft.Extensions.DependencyInjection;
using Microsoft.UI.Xaml;
using Microsoft.UI.Xaml.Controls;
using Microsoft.UI.Xaml.Controls.Primitives;
using Microsoft.UI.Xaml.Data;
using Microsoft.UI.Xaml.Input;
using Microsoft.UI.Xaml.Media;
using Microsoft.UI.Xaml.Navigation;
using ShiroProcessReporter.Services;
using ShiroProcessReporter.Components;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Runtime.InteropServices.WindowsRuntime;
using Windows.Foundation;
using Windows.Foundation.Collections;
using CommunityToolkit.Mvvm.Input;
using CommunityToolkit.Mvvm.ComponentModel;
using ShiroProcessReporter.Helper;

namespace ShiroProcessReporter.Views
{
[ObservableObject]
public sealed partial class EndpointSettings : Page
{
private readonly ReportService \_reportService;

        public EndpointSettings()
        {
            this.InitializeComponent();
            _reportService = App.ServiceProvider!.GetService<ReportService>()!;
        }

        [RelayCommand]
        private void UpdateEndpoint()
        {
            var context = EndpointEditDialog.DataContext as DataContextWrapper<string>;
            _reportService!.Endpoint = context.Value;
        }

        [RelayCommand]
        private void UpdateApiKey()
        {
            var context = EndpointEditDialog.DataContext as DataContextWrapper<string>;
            _reportService!.ApiKey = context.Value;
        }

        private async void EditEndpoint(object sender, RoutedEventArgs e)
        {
            EndpointEditDialog.DataContext = new DataContextWrapper<string>(_reportService!.Endpoint);
            EndpointEditDialog.Title = "Edit Endpoint";
            EndpointEditDialog.PrimaryButtonCommand = UpdateEndpointCommand;
            await EndpointEditDialog.ShowAsync();
        }

        private async void EditApiKey(object sender, RoutedEventArgs e)
        {
            EndpointEditDialog.DataContext = new DataContextWrapper<string>(_reportService!.ApiKey);
            EndpointEditDialog.Title = "Edit API Key";
            EndpointEditDialog.PrimaryButtonCommand = UpdateApiKeyCommand;
            await EndpointEditDialog.ShowAsync();
        }
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Views/FilterRuleSettings.xaml
================================================

<?xml version="1.0" encoding="utf-8"?>

<Page
    x:Class="ShiroProcessReporter.Views.FilterRuleSettings"
    xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
    xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"
    xmlns:local="using:ShiroProcessReporter.Views"
    xmlns:d="http://schemas.microsoft.com/expression/blend/2008"
    xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006"
    xmlns:tkconverters="using:CommunityToolkit.WinUI.Converters"
    xmlns:models="using:ShiroProcessReporter.Models"
    xmlns:ui="using:CommunityToolkit.WinUI"
    mc:Ignorable="d"
    Background="Transparent">

    <Page.Resources>
        <tkconverters:BoolToVisibilityConverter
            x:Key="BoolToInvertedVisibilityConverter"
            FalseValue="Visible"
            TrueValue="Collapsed" />
    </Page.Resources>

    <Grid HorizontalAlignment="Stretch" VerticalAlignment="Stretch" Padding="20" RowSpacing="8">
        <Grid.RowDefinitions>
            <RowDefinition Height="Auto"/>
            <RowDefinition Height="*" />
        </Grid.RowDefinitions>

        <ContentDialog
            x:Name="FilterRuleDialog"
            PrimaryButtonText="Save"
            SecondaryButtonText="Cancel"
            DefaultButton="Primary"
            IsPrimaryButtonEnabled="{Binding FilterRuleVaild, Mode=TwoWay}">
            <ContentDialog.DataContext>
                <models:FilterRule />
            </ContentDialog.DataContext>
            <StackPanel Spacing="8">
                <TextBlock Text="Pattern" />
                <TextBox
                    IsSpellCheckEnabled="False"
                    Text="{Binding Original, Mode=TwoWay, UpdateSourceTrigger=PropertyChanged}">
                    <TextBox.Description>
                        <StackPanel
                            Margin="0,4"
                            Orientation="Horizontal"
                            Spacing="4"
                            Visibility="{Binding IsOriginalVaild, Converter={StaticResource BoolToInvertedVisibilityConverter}}">
                            <FontIcon
                                Margin="0,0,0,0"
                                AutomationProperties.AccessibilityView="Raw"
                                FontSize="14"
                                Foreground="{ThemeResource SystemFillColorCautionBrush}"
                                Glyph="&#xE7BA;" />
                            <TextBlock Text="Pattern must be a valid regex expression" />
                        </StackPanel>
                    </TextBox.Description>
                </TextBox>
            </StackPanel>
        </ContentDialog>

        <ContentDialog
            x:Name="DeleteDialog"
            PrimaryButtonText="Sure"
            SecondaryButtonText="Cancel"
            PrimaryButtonCommand="{x:Bind DeleteCommand}"
            PrimaryButtonStyle="{StaticResource AccentButtonStyle}">
            <TextBlock Text="Are you sure delete this rule?" />
        </ContentDialog>

        <StackPanel HorizontalAlignment="Right" Grid.Row="0">
            <Button Click="{x:Bind OpenNewDialogAsync}">
                <StackPanel Orientation="Horizontal" Spacing="8">
                    <FontIcon
                FontSize="16"
                Foreground="{ThemeResource AccentTextFillColorPrimaryBrush}"
                Glyph="&#xe710;" />
                    <TextBlock Text="Add new" />
                </StackPanel>
                <Button.KeyboardAccelerators>
                    <KeyboardAccelerator Key="N" Modifiers="Control" />
                </Button.KeyboardAccelerators>
            </Button>
        </StackPanel>
        <ScrollView Grid.Row="1">
            <ListView
                x:Name="FilterRuleListView"
                Background="{ThemeResource LayerFillColorDefaultBrush}"
                BorderBrush="{ThemeResource CardStrokeColorDefaultBrush}"
                BorderThickness="1"
                CornerRadius="{StaticResource OverlayCornerRadius}"
                IsItemClickEnabled="True"
                ItemClick="ListView_ItemClick"
                ItemsSource="{x:Bind FilterRules, Mode=TwoWay}"
                RightTapped="ListView_RightTapped"
                SelectedItem="{x:Bind Selected, Mode=TwoWay}">
                <ListView.Header>
                    <Grid Background="Transparent" ColumnSpacing="8" HorizontalAlignment="Stretch" VerticalAlignment="Center" Margin="16,8">
                        <Grid.ColumnDefinitions>
                            <ColumnDefinition Width="*" />
                            <ColumnDefinition Width="64" />
                        </Grid.ColumnDefinitions>
                        <TextBlock Text="Pattern" Grid.Column="0" TextAlignment="Center" FontWeight="Bold" />
                    </Grid>
                </ListView.Header>
                <ListView.ContextFlyout>
                    <MenuFlyout>
                        <MenuFlyoutItem Text="Edit" Icon="Edit" Click="Edit_Click"/>
                        <MenuFlyoutItem Text="Duplicate" Click="Duplicate_Click">
                            <MenuFlyoutItem.Icon>
                                <FontIcon Glyph="&#xF413;" />
                            </MenuFlyoutItem.Icon>
                        </MenuFlyoutItem>
                        <MenuFlyoutSeparator />
                        <MenuFlyoutItem Text="MoveUp" Click="ReorderButtonUp_Click">
                            <MenuFlyoutItem.Icon>
                                <FontIcon Glyph="&#xE74A;" />
                            </MenuFlyoutItem.Icon>
                        </MenuFlyoutItem>
                        <MenuFlyoutItem Text="MoveDown" Click="ReorderButtonDown_Click">
                            <MenuFlyoutItem.Icon>
                                <FontIcon Glyph="&#xE74B;" />
                            </MenuFlyoutItem.Icon>
                        </MenuFlyoutItem>
                        <MenuFlyoutSeparator />
                        <MenuFlyoutItem Text="Delete" Icon="Delete" Click="Delete_Click" />
                    </MenuFlyout>
                </ListView.ContextFlyout>
                <ListView.ItemTemplate>
                    <DataTemplate x:DataType="models:FilterRule">
                        <Grid Background="Transparent" ColumnSpacing="8" HorizontalAlignment="Stretch" VerticalAlignment="Center">
                            <Grid.ColumnDefinitions>
                                <ColumnDefinition Width="*" />
                                <ColumnDefinition Width="64" />
                            </Grid.ColumnDefinitions>
                            <TextBlock Grid.Column="0" Text="{Binding Original, Mode=OneWay}" VerticalAlignment="Center" TextAlignment="Center" />
                            <Button
                                Grid.Column="3"
                                Height="32"
                                Content="{ui:FontIcon Glyph=&#xE74D;,FontSize=16}"
                                Background="Transparent"
                                BorderBrush="Transparent"
                                Click="Delete_Click"
                                CommandParameter="{x:Bind (models:FilterRule)}"
                                GotFocus="Item_GotFocus"/>
                        </Grid>
                    </DataTemplate>
                </ListView.ItemTemplate>
            </ListView>
        </ScrollView>
    </Grid>

</Page>

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Views/FilterRuleSettings.xaml.cs
================================================
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Runtime.InteropServices.WindowsRuntime;
using Windows.Foundation;
using Windows.Foundation.Collections;
using Microsoft.UI.Xaml;
using Microsoft.UI.Xaml.Controls;
using Microsoft.UI.Xaml.Controls.Primitives;
using Microsoft.UI.Xaml.Data;
using Microsoft.UI.Xaml.Input;
using Microsoft.UI.Xaml.Media;
using Microsoft.UI.Xaml.Navigation;
using CommunityToolkit.Mvvm.ComponentModel;
using System.Collections.ObjectModel;
using Microsoft.Extensions.DependencyInjection;
using ShiroProcessReporter.Services;
using CommunityToolkit.Mvvm.Input;
using System.Threading.Tasks;
using ShiroProcessReporter.Models;

namespace ShiroProcessReporter.Views
{
[ObservableObject]
public sealed partial class FilterRuleSettings : Page
{
private readonly ReportService \_reportService;

        [ObservableProperty]
        private FilterRule? _selected;

        [ObservableProperty]
        private ObservableCollection<FilterRule> _filterRules;

        public FilterRuleSettings()
        {
            this.InitializeComponent();
            this._reportService = App.ServiceProvider!.GetService<ReportService>()!;
            this._filterRules = new ObservableCollection<FilterRule>(_reportService.FilterRules);
        }

        [RelayCommand]
        private void Add()
        {
            _filterRules.Add(FilterRuleDialog.DataContext as FilterRule);
            _reportService!.FilterRules = [.. _filterRules];
        }

        [RelayCommand]
        private void Update()
        {
            _filterRules[FilterRuleListView.SelectedIndex] = FilterRuleDialog.DataContext as FilterRule;
            _reportService!.FilterRules = [.. _filterRules];
        }

        [RelayCommand]
        private void Delete()
        {
            _filterRules.RemoveAt(FilterRuleListView.SelectedIndex);
            _reportService!.FilterRules = [.. _filterRules];
        }

        private async Task OpenNewDialogAsync(object sender, object e)
        {
            await ShowAddDialogAsync();
        }

        private void ListView_RightTapped(object sender, RightTappedRoutedEventArgs e)
        {
            var rule = (e.OriginalSource as FrameworkElement).DataContext as FilterRule;
            Selected = rule;
        }

        private async void ListView_ItemClick(object sender, ItemClickEventArgs e)
        {
            FilterRule rule = e.ClickedItem as FilterRule;
            Selected = rule;
            await ShowEditDialogAsync(rule);
        }

        private async void Delete_Click(object sender, RoutedEventArgs e)
        {
            if (FilterRuleListView.SelectedItem is FilterRule rule)
            {
                Selected = rule;
                DeleteDialog.Title = $"Filter \"{rule.Original}\"";
                await DeleteDialog.ShowAsync();
            }
        }

        private async void Edit_Click(object sender, RoutedEventArgs e)
        {
            if (FilterRuleListView.SelectedItem is FilterRule rule)
            {
                await ShowEditDialogAsync(rule);
            }
        }

        private async void Duplicate_Click(object sender, RoutedEventArgs e)
        {
            if (FilterRuleListView.SelectedItem is FilterRule rule)
            {
                await ShowAddDialogAsync(rule);
            }
        }

        private void ReorderButtonUp_Click(object sender, RoutedEventArgs e)
        {
            if (FilterRuleListView.SelectedItem is FilterRule rule)
            {
                var index = FilterRules.IndexOf(rule);
                if (index > 0)
                {
                    FilterRules.Move(index, index - 1);
                    _reportService!.FilterRules = [.. _filterRules];

                }
            }
        }

        private void ReorderButtonDown_Click(object sender, RoutedEventArgs e)
        {
            if (FilterRuleListView.SelectedItem is FilterRule rule)
            {
                var index = FilterRules.IndexOf(rule);
                if (index < FilterRules.Count - 1)
                {
                    FilterRules.Move(index, index + 1);
                    _reportService!.FilterRules = [.. _filterRules];
                }
            }
        }

        private void Item_GotFocus(object sender, RoutedEventArgs e)
        {
            var element = sender as FrameworkElement;
            var rule = element.DataContext as FilterRule;

            if (rule is not null)
            {
                Selected = rule;
            }
            else if (FilterRuleListView.SelectedItem == null && FilterRuleListView.Items.Count > 0)
            {
                FilterRuleListView.SelectedItem = 0;
            }
        }

        private async Task ShowEditDialogAsync(FilterRule rule)
        {
            if (Selected is null)
            {
                return;
            }
            FilterRuleDialog.Title = "Edit filter rule";
            FilterRuleDialog.PrimaryButtonCommand = UpdateCommand;
            var clone = Selected.Clone();
            FilterRuleDialog.DataContext = clone;
            await FilterRuleDialog.ShowAsync();
        }

        private async Task ShowAddDialogAsync(FilterRule? template = null)
        {
            FilterRuleDialog.Title = "Add new filter rule";
            FilterRuleDialog.PrimaryButtonCommand = AddCommand;
            if (template is not null)
            {
                var clone = template.Clone();
                FilterRuleDialog.DataContext = clone;
            }

            await FilterRuleDialog.ShowAsync();
        }
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Views/LogViewer.xaml
================================================

<?xml version="1.0" encoding="utf-8"?>

<Page
    x:Class="ShiroProcessReporter.Views.LogViewer"
    xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
    xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"
    xmlns:local="using:ShiroProcessReporter.Views"
    xmlns:d="http://schemas.microsoft.com/expression/blend/2008"
    xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006"
    xmlns:ui="using:CommunityToolkit.WinUI"
    xmlns:helper="using:ShiroProcessReporter.Helper"
    xmlns:models="using:ShiroProcessReporter.Models"
    xmlns:converters="using:ShiroProcessReporter.Converters"
    d:DataContext="{d:DesignInstance Type=local:LogViewer}"
    mc:Ignorable="d"
    Background="Transparent">

    <Page.Resources>
        <converters:LogLevelToColorConverter x:Key="LogLevelToColorConverter" />
    </Page.Resources>

    <Grid HorizontalAlignment="Stretch" VerticalAlignment="Stretch" Padding="20" RowSpacing="8">
        <Grid.RowDefinitions>
            <RowDefinition Height="Auto"/>
            <RowDefinition Height="*" />
        </Grid.RowDefinitions>
        <StackPanel Orientation="Horizontal" HorizontalAlignment="Right" Spacing="4" Grid.Row="0">
            <ToggleButton
                Height="32"
                Content="{ui:FontIcon Glyph=&#xE74B;,FontSize=16}"
                Background="Transparent"
                BorderBrush="Transparent"
                Checked="FollowToggleButton_Checked"
                Unchecked="FollowToggleButton_Unchecked">
                <ToolTipService.ToolTip>
                    <ToolTip Content="Follow" />
                </ToolTipService.ToolTip>
            </ToggleButton>
            <Button
                Height="32"
                Content="{ui:FontIcon Glyph=&#xE74D;,FontSize=16}"
                Background="Transparent"
                BorderBrush="Transparent">
                <ToolTipService.ToolTip>
                    <ToolTip Content="Clean all logs" />
                </ToolTipService.ToolTip>
            </Button>
        </StackPanel>
        <ListView
            Grid.Row="1"
            x:Name="LogListView">
            <ListView.ItemsPanel>
                <ItemsPanelTemplate>
                    <ItemsStackPanel />
                </ItemsPanelTemplate>
            </ListView.ItemsPanel>
            <ListView.ItemTemplate>
                <DataTemplate x:DataType="models:LogEntry">
                    <TextBlock
                        FontFamily="'Cascadia Mono', monospace"
                        Text="{Binding Message}"
                        Foreground="{Binding LogLevel, Converter={StaticResource LogLevelToColorConverter}}"/>
                </DataTemplate>
            </ListView.ItemTemplate>
        </ListView>
    </Grid>

</Page>

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Views/LogViewer.xaml.cs
================================================
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Runtime.InteropServices.WindowsRuntime;
using Windows.Foundation;
using Windows.Foundation.Collections;
using Microsoft.UI.Xaml;
using Microsoft.UI.Xaml.Controls;
using Microsoft.UI.Xaml.Controls.Primitives;
using Microsoft.UI.Xaml.Data;
using Microsoft.UI.Xaml.Input;
using Microsoft.UI.Xaml.Media;
using Microsoft.UI.Xaml.Navigation;
using ShiroProcessReporter.Models;
using System.Collections.ObjectModel;
using CommunityToolkit.Mvvm.ComponentModel;
using ShiroProcessReporter.Helper;
using CommunityToolkit.Mvvm.Input;
using Microsoft.Extensions.Logging.Abstractions;

namespace ShiroProcessReporter.Views
{
[ObservableObject]
public sealed partial class LogViewer : Page
{
[ObservableProperty]
private bool \_isFollow = false;

        public LogViewer()
        {
            this.InitializeComponent();
        }

        protected override void OnNavigatedTo(NavigationEventArgs e)
        {
            base.OnNavigatedTo(e);
            GlobalState.Instance.LogViewDispatcherQueue = DispatcherQueue;
            GlobalState.Instance.Logs.CollectionChanged += Logs_CollectionChanged;
            LogListView.SetBinding(
                ItemsControl.ItemsSourceProperty,
                new Binding() { Source = GlobalState.Instance.Logs, Mode = BindingMode.OneWay});
        }

        protected override void OnNavigatedFrom(NavigationEventArgs e)
        {
            LogListView.ClearValue(ItemsControl.ItemsSourceProperty);
            GlobalState.Instance.Logs.CollectionChanged -= Logs_CollectionChanged;
            GlobalState.Instance.LogViewDispatcherQueue = null;
            base.OnNavigatedFrom(e);
        }

        private void Logs_CollectionChanged(object sender, System.Collections.Specialized.NotifyCollectionChangedEventArgs e)
        {
            if (IsFollow)
            {
                ScrollToBottom();
            }
        }

        private void FollowToggleButton_Checked(object sender, RoutedEventArgs e)
        {
            IsFollow = true;
            ScrollToBottom();
        }

        private void FollowToggleButton_Unchecked(object sender, RoutedEventArgs e)
        {
            IsFollow = false;
        }

        private void ScrollToBottom()
        {
            var scrollViewer = GetScrollViewer(LogListView);
            if (scrollViewer is not null)
            {
                DispatcherQueue.TryEnqueue(() =>
                {
                    var scrollableHeight = scrollViewer.ScrollableHeight;
                    scrollViewer.ChangeView(null, scrollableHeight, null, true);
                });
            }
        }

        // ¸¨Öú·½·¨£ºµÝ¹é²éÕÒListViewÖÐµÄScrollViewer
        private ScrollViewer GetScrollViewer(DependencyObject dependencyObject)
        {
            if (dependencyObject is ScrollViewer)
            {
                return dependencyObject as ScrollViewer;
            }

            for (int i = 0; i < VisualTreeHelper.GetChildrenCount(dependencyObject); i++)
            {
                var child = VisualTreeHelper.GetChild(dependencyObject, i);
                var result = GetScrollViewer(child);
                if (result != null)
                {
                    return result;
                }
            }
            return null;
        }

        [RelayCommand]
        private void ClearLogs()
        {
            GlobalState.Instance.Logs.Clear();
        }
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Views/ReplaceRuleSettings.xaml
================================================

<?xml version="1.0" encoding="utf-8"?>

<Page
    x:Class="ShiroProcessReporter.Views.ReplaceRuleSettings"
    xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
    xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"
    xmlns:local="using:ShiroProcessReporter.Views"
    xmlns:d="http://schemas.microsoft.com/expression/blend/2008"
    xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006"
    xmlns:tkconverters="using:CommunityToolkit.WinUI.Converters"
    xmlns:models="using:ShiroProcessReporter.Models"
    xmlns:ui="using:CommunityToolkit.WinUI"
    mc:Ignorable="d"
    Background="Transparent">

    <Page.Resources>
        <tkconverters:BoolToVisibilityConverter
            x:Key="BoolToInvertedVisibilityConverter"
            FalseValue="Visible"
            TrueValue="Collapsed" />
    </Page.Resources>

    <Grid HorizontalAlignment="Stretch" VerticalAlignment="Stretch" Padding="20" RowSpacing="8">
        <Grid.RowDefinitions>
            <RowDefinition Height="Auto"/>
            <RowDefinition Height="*" />
        </Grid.RowDefinitions>

        <ContentDialog
            x:Name="ReplaceRuleDialog"
            PrimaryButtonText="Save"
            SecondaryButtonText="Cancel"
            DefaultButton="Primary"
            IsPrimaryButtonEnabled="{Binding ReplaceRuleVaild, Mode=TwoWay}">
            <ContentDialog.DataContext>
                <models:ReplaceRule />
            </ContentDialog.DataContext>
            <StackPanel Spacing="8">
                <TextBlock Text="Pattern" />
                <TextBox
                    IsSpellCheckEnabled="False"
                    Text="{Binding Original, Mode=TwoWay, UpdateSourceTrigger=PropertyChanged}">
                    <TextBox.Description>
                        <StackPanel
                            Margin="0,4"
                            Orientation="Horizontal"
                            Spacing="4"
                            Visibility="{Binding IsOriginalVaild, Converter={StaticResource BoolToInvertedVisibilityConverter}}">
                            <FontIcon
                                    Margin="0,0,0,0"
                                    AutomationProperties.AccessibilityView="Raw"
                                    FontSize="14"
                                    Foreground="{ThemeResource SystemFillColorCautionBrush}"
                                    Glyph="&#xE7BA;" />
                            <TextBlock Text="Pattern must be a valid regex expression" />
                        </StackPanel>
                    </TextBox.Description>
                </TextBox>
                <TextBlock Text="Replace to" />
                <TextBox
                    IsSpellCheckEnabled="False"
                    Text="{Binding Replacement, Mode=TwoWay, UpdateSourceTrigger=PropertyChanged}" />
            </StackPanel>
        </ContentDialog>

        <ContentDialog
            x:Name="DeleteDialog"
            PrimaryButtonText="Sure"
            SecondaryButtonText="Cancel"
            PrimaryButtonCommand="{x:Bind DeleteCommand}"
            PrimaryButtonStyle="{StaticResource AccentButtonStyle}">
            <TextBlock Text="Are you sure delete this rule?" />
        </ContentDialog>

        <StackPanel HorizontalAlignment="Right" Grid.Row="0">
            <Button Click="{x:Bind OpenNewDialogAsync}">
                <StackPanel Orientation="Horizontal" Spacing="8">
                    <FontIcon
                        FontSize="16"
                        Foreground="{ThemeResource AccentTextFillColorPrimaryBrush}"
                        Glyph="&#xe710;" />
                    <TextBlock Text="Add new" />
                </StackPanel>
                <Button.KeyboardAccelerators>
                    <KeyboardAccelerator Key="N" Modifiers="Control" />
                </Button.KeyboardAccelerators>
            </Button>
        </StackPanel>
        <ScrollView Grid.Row="1">
            <ListView
                x:Name="ReplaceRuleListView"
                Background="{ThemeResource LayerFillColorDefaultBrush}"
                BorderBrush="{ThemeResource CardStrokeColorDefaultBrush}"
                BorderThickness="1"
                CornerRadius="{StaticResource OverlayCornerRadius}"
                IsItemClickEnabled="True"
                ItemClick="ListView_ItemClick"
                ItemsSource="{x:Bind ReplaceRules, Mode=TwoWay}"
                RightTapped="ListView_RightTapped"
                SelectedItem="{x:Bind Selected, Mode=TwoWay}">
                <ListView.Header>
                    <Grid Background="Transparent" ColumnSpacing="8" HorizontalAlignment="Stretch" VerticalAlignment="Center" Margin="16,8">
                        <Grid.ColumnDefinitions>
                            <ColumnDefinition Width="*" />
                            <ColumnDefinition Width="20" />
                            <ColumnDefinition Width="*" />
                            <ColumnDefinition Width="64" />
                        </Grid.ColumnDefinitions>
                        <TextBlock Text="Pattern" Grid.Column="0" TextAlignment="Center" FontWeight="Bold" />
                        <TextBlock Text="Replace to" Grid.Column="2" TextAlignment="Center" FontWeight="Bold" />
                    </Grid>
                </ListView.Header>
                <ListView.ContextFlyout>
                    <MenuFlyout>
                        <MenuFlyoutItem Text="Edit" Icon="Edit" Click="Edit_Click"/>
                        <MenuFlyoutItem Text="Duplicate" Click="Duplicate_Click">
                            <MenuFlyoutItem.Icon>
                                <FontIcon Glyph="&#xF413;" />
                            </MenuFlyoutItem.Icon>
                        </MenuFlyoutItem>
                        <MenuFlyoutSeparator />
                        <MenuFlyoutItem Text="MoveUp" Click="ReorderButtonUp_Click">
                            <MenuFlyoutItem.Icon>
                                <FontIcon Glyph="&#xE74A;" />
                            </MenuFlyoutItem.Icon>
                        </MenuFlyoutItem>
                        <MenuFlyoutItem Text="MoveDown" Click="ReorderButtonDown_Click">
                            <MenuFlyoutItem.Icon>
                                <FontIcon Glyph="&#xE74B;" />
                            </MenuFlyoutItem.Icon>
                        </MenuFlyoutItem>
                        <MenuFlyoutSeparator />
                        <MenuFlyoutItem Text="Delete" Icon="Delete" Click="Delete_Click" />
                    </MenuFlyout>
                </ListView.ContextFlyout>
                <ListView.ItemTemplate>
                    <DataTemplate x:DataType="models:ReplaceRule">
                        <Grid Background="Transparent" ColumnSpacing="8" HorizontalAlignment="Stretch" VerticalAlignment="Center">
                            <Grid.ColumnDefinitions>
                                <ColumnDefinition Width="*" />
                                <ColumnDefinition Width="20" />
                                <ColumnDefinition Width="*" />
                                <ColumnDefinition Width="64" />
                            </Grid.ColumnDefinitions>
                            <TextBlock Grid.Column="0" Text="{Binding Original, Mode=OneWay}" VerticalAlignment="Center" TextAlignment="Center" />
                            <Viewbox Grid.Column="1" Height="12" Width="12" Margin="0,4,0,0">
                                <SymbolIcon Symbol="Forward" />
                            </Viewbox>
                            <TextBlock Grid.Column="2" Text="{Binding Replacement, Mode=OneWay}" VerticalAlignment="Center" TextAlignment="Center" />
                            <Button
                                Grid.Column="3"
                                Height="32"
                                Content="{ui:FontIcon Glyph=&#xE74D;,FontSize=16}"
                                Background="Transparent"
                                BorderBrush="Transparent"
                                Click="Delete_Click"
                                CommandParameter="{x:Bind (models:ReplaceRule)}"
                                GotFocus="Item_GotFocus"/>
                        </Grid>
                    </DataTemplate>
                </ListView.ItemTemplate>
            </ListView>
        </ScrollView>
    </Grid>

</Page>

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter/Views/ReplaceRuleSettings.xaml.cs
================================================
using CommunityToolkit.Mvvm.ComponentModel;
using CommunityToolkit.Mvvm.Input;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.UI.Xaml;
using Microsoft.UI.Xaml.Controls;
using Microsoft.UI.Xaml.Controls.Primitives;
using Microsoft.UI.Xaml.Data;
using Microsoft.UI.Xaml.Input;
using Microsoft.UI.Xaml.Media;
using Microsoft.UI.Xaml.Navigation;
using ShiroProcessReporter.Components;
using ShiroProcessReporter.Models;
using ShiroProcessReporter.Services;
using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.IO;
using System.Linq;
using System.Runtime.InteropServices.WindowsRuntime;
using System.Threading.Tasks;
using System.Windows.Input;
using Windows.Foundation;
using Windows.Foundation.Collections;

namespace ShiroProcessReporter.Views
{
[ObservableObject]
public sealed partial class ReplaceRuleSettings : Page
{
private readonly ReportService \_reportService;

        [ObservableProperty]
        private ReplaceRule? _selected;

        [ObservableProperty]
        private ObservableCollection<ReplaceRule> _replaceRules;

        public ReplaceRuleSettings()
        {
            this.InitializeComponent();
            this._reportService = App.ServiceProvider!.GetService<ReportService>()!;
            this._replaceRules = new ObservableCollection<ReplaceRule>(_reportService.ReplaceRules);
        }

        [RelayCommand]
        private void Add()
        {
            _replaceRules.Add(ReplaceRuleDialog.DataContext as ReplaceRule);
            _reportService!.ReplaceRules = [.. _replaceRules];
        }

        [RelayCommand]
        private void Update()
        {
            _replaceRules[ReplaceRuleListView.SelectedIndex] = ReplaceRuleDialog.DataContext as ReplaceRule;
            _reportService!.ReplaceRules = [.. _replaceRules];
        }

        [RelayCommand]
        private void Delete()
        {
            _replaceRules.RemoveAt(ReplaceRuleListView.SelectedIndex);
            _reportService!.ReplaceRules = [.. _replaceRules];
        }

        private async Task OpenNewDialogAsync(object sender, object e)
        {
            await ShowAddDialogAsync();
        }

        private void ListView_RightTapped(object sender, RightTappedRoutedEventArgs e)
        {
            var rule = (e.OriginalSource as FrameworkElement).DataContext as ReplaceRule;
            Selected = rule;
        }

        private async void ListView_ItemClick(object sender, ItemClickEventArgs e)
        {
            ReplaceRule rule = e.ClickedItem as ReplaceRule;
            Selected = rule;
            await ShowEditDialogAsync(rule);
        }

        private async void Delete_Click(object sender, RoutedEventArgs e)
        {
            if (ReplaceRuleListView.SelectedItem is ReplaceRule rule)
            {
                Selected = rule;
                DeleteDialog.Title = $"Replace \"{rule.Original}\" to {rule.Replacement}";
                await DeleteDialog.ShowAsync();
            }
        }

        private async void Edit_Click(object sender, RoutedEventArgs e)
        {
            if (ReplaceRuleListView.SelectedItem is ReplaceRule rule)
            {
                await ShowEditDialogAsync(rule);
            }
        }

        private async void Duplicate_Click(object sender, RoutedEventArgs e)
        {
            if (ReplaceRuleListView.SelectedItem is ReplaceRule rule)
            {
                await ShowAddDialogAsync(rule);
            }
        }

        private void ReorderButtonUp_Click(object sender, RoutedEventArgs e)
        {
            if (ReplaceRuleListView.SelectedItem is ReplaceRule rule)
            {
                var index = ReplaceRules.IndexOf(rule);
                if (index > 0)
                {
                    ReplaceRules.Move(index, index - 1);
                    _reportService!.ReplaceRules = [.. _replaceRules];
                }
            }
        }

        private void ReorderButtonDown_Click(object sender, RoutedEventArgs e)
        {
            if (ReplaceRuleListView.SelectedItem is ReplaceRule rule)
            {
                var index = ReplaceRules.IndexOf(rule);
                if (index < ReplaceRules.Count - 1)
                {
                    ReplaceRules.Move(index, index + 1);
                    _reportService!.ReplaceRules = [.. _replaceRules];
                }
            }
        }

        private void Item_GotFocus(object sender, RoutedEventArgs e)
        {
            var element = sender as FrameworkElement;
            var rule = element.DataContext as ReplaceRule;

            if (rule is not null)
            {
                Selected = rule;
            }
            else if (ReplaceRuleListView.SelectedItem == null && ReplaceRuleListView.Items.Count > 0)
            {
                ReplaceRuleListView.SelectedItem = 0;
            }
        }

        private async Task ShowEditDialogAsync(ReplaceRule rule)
        {
            if (Selected is null)
            {
                return;
            }
            ReplaceRuleDialog.Title = "Edit replace rule";
            ReplaceRuleDialog.PrimaryButtonCommand = UpdateCommand;
            var clone = Selected.Clone();
            ReplaceRuleDialog.DataContext = clone;
            await ReplaceRuleDialog.ShowAsync();
        }

        private async Task ShowAddDialogAsync(ReplaceRule? template = null)
        {
            ReplaceRuleDialog.Title = "Add new replace rule";
            ReplaceRuleDialog.PrimaryButtonCommand = AddCommand;
            if (template is not null)
            {
                var clone = template.Clone();
                ReplaceRuleDialog.DataContext = clone;
            }

            await ReplaceRuleDialog.ShowAsync();
        }
    }

}

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter (Package)/Package.appxmanifest
================================================
﻿<?xml version="1.0" encoding="utf-8"?>

<Package
  xmlns="http://schemas.microsoft.com/appx/manifest/foundation/windows10"
  xmlns:mp="http://schemas.microsoft.com/appx/2014/phone/manifest"
  xmlns:uap="http://schemas.microsoft.com/appx/manifest/uap/windows10"
  xmlns:rescap="http://schemas.microsoft.com/appx/manifest/foundation/windows10/restrictedcapabilities"
  IgnorableNamespaces="uap rescap">

<Identity
    Name="f8f2aa4b-ae72-4e86-b26e-e0e955af5ae2"
    Publisher="CN=ChingC"
    Version="2.0.0.0" />

<mp:PhoneIdentity PhoneProductId="f8f2aa4b-ae72-4e86-b26e-e0e955af5ae2" PhonePublisherId="00000000-0000-0000-0000-000000000000"/>

  <Properties>
    <DisplayName>ShiroProcessReporter</DisplayName>
    <PublisherDisplayName>ChingC</PublisherDisplayName>
    <Logo>Images\StoreLogo.png</Logo>
  </Properties>

  <Dependencies>
    <TargetDeviceFamily Name="Windows.Universal" MinVersion="10.0.17763.0" MaxVersionTested="10.0.19041.0" />
    <TargetDeviceFamily Name="Windows.Desktop" MinVersion="10.0.17763.0" MaxVersionTested="10.0.19041.0" />
  </Dependencies>

  <Resources>
    <Resource Language="x-generate"/>
  </Resources>

  <Applications>
    <Application Id="App"
      Executable="$targetnametoken$.exe"
      EntryPoint="$targetentrypoint$">
      <uap:VisualElements
        DisplayName="ShiroProcessReporter"
        Description="ShiroProcessReporter"
        BackgroundColor="transparent"
        Square150x150Logo="Images\Square150x150Logo.png"
        Square44x44Logo="Images\Square44x44Logo.png">
        <uap:DefaultTile Wide310x150Logo="Images\Wide310x150Logo.png"  Square71x71Logo="Images\SmallTile.png" Square310x310Logo="Images\LargeTile.png"/>
        <uap:SplashScreen Image="Images\SplashScreen.png" />
      </uap:VisualElements>
    </Application>
  </Applications>

  <Capabilities>
    <rescap:Capability Name="runFullTrust" />
  </Capabilities>
</Package>

================================================
FILE: ShiroProcessReporter/ShiroProcessReporter (Package)/ShiroProcessReporter (Package).wapproj
================================================

<?xml version="1.0" encoding="utf-8"?>
<Project ToolsVersion="15.0" DefaultTargets="Build" xmlns="http://schemas.microsoft.com/developer/msbuild/2003">
  <PropertyGroup Condition="'$(VisualStudioVersion)' == '' or '$(VisualStudioVersion)' &lt; '15.0'">
    <VisualStudioVersion>15.0</VisualStudioVersion>
  </PropertyGroup>
  <ItemGroup Label="ProjectConfigurations">
    <ProjectConfiguration Include="Debug|x86">
      <Configuration>Debug</Configuration>
      <Platform>x86</Platform>
    </ProjectConfiguration>
    <ProjectConfiguration Include="Release|x86">
      <Configuration>Release</Configuration>
      <Platform>x86</Platform>
    </ProjectConfiguration>
    <ProjectConfiguration Include="Debug|x64">
      <Configuration>Debug</Configuration>
      <Platform>x64</Platform>
    </ProjectConfiguration>
    <ProjectConfiguration Include="Release|x64">
      <Configuration>Release</Configuration>
      <Platform>x64</Platform>
    </ProjectConfiguration>
    <ProjectConfiguration Include="Debug|ARM64">
      <Configuration>Debug</Configuration>
      <Platform>ARM64</Platform>
    </ProjectConfiguration>
    <ProjectConfiguration Include="Release|ARM64">
      <Configuration>Release</Configuration>
      <Platform>ARM64</Platform>
    </ProjectConfiguration>
  </ItemGroup>
  <PropertyGroup>
    <WapProjPath Condition="'$(WapProjPath)'==''">$(MSBuildExtensionsPath)\Microsoft\DesktopBridge\</WapProjPath>
    <PathToXAMLWinRTImplementations>ShiroProcessReporter\</PathToXAMLWinRTImplementations>
  </PropertyGroup>
  <Import Project="$(WapProjPath)\Microsoft.DesktopBridge.props" />
  <PropertyGroup>
    <ProjectGuid>ac822e31-7c2d-4f32-884b-db0b5fdda414</ProjectGuid>
    <TargetPlatformVersion>10.0.22621.0</TargetPlatformVersion>
    <TargetPlatformMinVersion>10.0.17763.0</TargetPlatformMinVersion>
    <AssetTargetFallback>net6.0-windows$(TargetPlatformVersion);$(AssetTargetFallback)</AssetTargetFallback>
    <DefaultLanguage>en-US</DefaultLanguage>
    <AppxPackageSigningEnabled>True</AppxPackageSigningEnabled>
    <EntryPointProjectUniqueName>..\ShiroProcessReporter\ShiroProcessReporter.csproj</EntryPointProjectUniqueName>
    <GenerateAppInstallerFile>False</GenerateAppInstallerFile>
    <AppxPackageSigningTimestampDigestAlgorithm>SHA256</AppxPackageSigningTimestampDigestAlgorithm>
    <AppxAutoIncrementPackageRevision>False</AppxAutoIncrementPackageRevision>
    <GenerateTestArtifacts>True</GenerateTestArtifacts>
    <AppxBundlePlatforms>x64</AppxBundlePlatforms>
    <HoursBetweenUpdateChecks>0</HoursBetweenUpdateChecks>
    <PackageCertificateThumbprint>3A2AA6497327DE69AC71FA94C7BD6990E56C545A</PackageCertificateThumbprint>
  </PropertyGroup>
  <PropertyGroup Condition="'$(Configuration)|$(Platform)'=='Release|x86'">
    <AppxBundle>Always</AppxBundle>
  </PropertyGroup>
  <PropertyGroup Condition="'$(Configuration)|$(Platform)'=='Debug|x86'">
    <AppxBundle>Always</AppxBundle>
  </PropertyGroup>
  <PropertyGroup Condition="'$(Configuration)|$(Platform)'=='Release|x64'">
    <AppxBundle>Always</AppxBundle>
  </PropertyGroup>
  <PropertyGroup Condition="'$(Configuration)|$(Platform)'=='Release|ARM64'">
    <AppxBundle>Always</AppxBundle>
  </PropertyGroup>
  <PropertyGroup Condition="'$(Configuration)|$(Platform)'=='Debug|ARM64'">
    <AppxBundle>Always</AppxBundle>
  </PropertyGroup>
  <PropertyGroup Condition="'$(Configuration)|$(Platform)'=='Debug|x64'">
    <AppxBundle>Always</AppxBundle>
  </PropertyGroup>
  <ItemGroup>
    <AppxManifest Include="Package.appxmanifest">
      <SubType>Designer</SubType>
    </AppxManifest>
  </ItemGroup>
  <ItemGroup>
    <Content Include="Images\LargeTile.scale-100.png" />
    <Content Include="Images\LargeTile.scale-125.png" />
    <Content Include="Images\LargeTile.scale-150.png" />
    <Content Include="Images\LargeTile.scale-200.png" />
    <Content Include="Images\LargeTile.scale-400.png" />
    <Content Include="Images\SmallTile.scale-100.png" />
    <Content Include="Images\SmallTile.scale-125.png" />
    <Content Include="Images\SmallTile.scale-150.png" />
    <Content Include="Images\SmallTile.scale-200.png" />
    <Content Include="Images\SmallTile.scale-400.png" />
    <Content Include="Images\SplashScreen.scale-100.png" />
    <Content Include="Images\SplashScreen.scale-125.png" />
    <Content Include="Images\SplashScreen.scale-150.png" />
    <Content Include="Images\SplashScreen.scale-200.png" />
    <Content Include="Images\LockScreenLogo.scale-200.png" />
    <Content Include="Images\SplashScreen.scale-400.png" />
    <Content Include="Images\Square150x150Logo.scale-100.png" />
    <Content Include="Images\Square150x150Logo.scale-125.png" />
    <Content Include="Images\Square150x150Logo.scale-150.png" />
    <Content Include="Images\Square150x150Logo.scale-200.png" />
    <Content Include="Images\Square150x150Logo.scale-400.png" />
    <Content Include="Images\Square44x44Logo.altform-lightunplated_targetsize-16.png" />
    <Content Include="Images\Square44x44Logo.altform-lightunplated_targetsize-24.png" />
    <Content Include="Images\Square44x44Logo.altform-lightunplated_targetsize-256.png" />
    <Content Include="Images\Square44x44Logo.altform-lightunplated_targetsize-32.png" />
    <Content Include="Images\Square44x44Logo.altform-lightunplated_targetsize-48.png" />
    <Content Include="Images\Square44x44Logo.altform-unplated_targetsize-16.png" />
    <Content Include="Images\Square44x44Logo.altform-unplated_targetsize-256.png" />
    <Content Include="Images\Square44x44Logo.altform-unplated_targetsize-32.png" />
    <Content Include="Images\Square44x44Logo.altform-unplated_targetsize-48.png" />
    <Content Include="Images\Square44x44Logo.scale-100.png" />
    <Content Include="Images\Square44x44Logo.scale-125.png" />
    <Content Include="Images\Square44x44Logo.scale-150.png" />
    <Content Include="Images\Square44x44Logo.scale-200.png" />
    <Content Include="Images\Square44x44Logo.scale-400.png" />
    <Content Include="Images\Square44x44Logo.targetsize-16.png" />
    <Content Include="Images\Square44x44Logo.targetsize-24.png" />
    <Content Include="Images\Square44x44Logo.targetsize-24_altform-unplated.png" />
    <Content Include="Images\Square44x44Logo.targetsize-256.png" />
    <Content Include="Images\Square44x44Logo.targetsize-32.png" />
    <Content Include="Images\Square44x44Logo.targetsize-48.png" />
    <Content Include="Images\StoreLogo.scale-100.png" />
    <Content Include="Images\StoreLogo.scale-125.png" />
    <Content Include="Images\StoreLogo.scale-150.png" />
    <Content Include="Images\StoreLogo.scale-200.png" />
    <Content Include="Images\StoreLogo.scale-400.png" />
    <Content Include="Images\Wide310x150Logo.scale-100.png" />
    <Content Include="Images\Wide310x150Logo.scale-125.png" />
    <Content Include="Images\Wide310x150Logo.scale-150.png" />
    <Content Include="Images\Wide310x150Logo.scale-200.png" />
    <Content Include="Images\Wide310x150Logo.scale-400.png" />
    <None Include="ShiroProcessReporter %28Package%29_TemporaryKey.pfx" />
  </ItemGroup>
  <ItemGroup>
    <ProjectReference Include="..\ShiroProcessReporter\ShiroProcessReporter.csproj">
      <SkipGetTargetFrameworkProperties>True</SkipGetTargetFrameworkProperties>
      <PublishProfile>Properties\PublishProfiles\win-$(Platform).pubxml</PublishProfile>
    </ProjectReference>
  </ItemGroup>
  <ItemGroup>
    <PackageReference Include="Microsoft.WindowsAppSDK" Version="1.5.240428000">
      <IncludeAssets>build</IncludeAssets>
    </PackageReference>
    <PackageReference Include="Microsoft.Windows.SDK.BuildTools" Version="10.0.22621.3233">
      <IncludeAssets>build</IncludeAssets>
    </PackageReference>
  </ItemGroup>
  <Import Project="$(WapProjPath)\Microsoft.DesktopBridge.targets" />
</Project>
