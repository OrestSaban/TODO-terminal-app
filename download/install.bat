@echo off

REM Define the version and platform
SET VERSION=v1.0
SET PLATFORM=windows-amd64.exe

REM Define the download URL
SET URL=https://github.com/yourusername/yourrepo/releases/download/%VERSION%/todo-%PLATFORM%

REM Download the binary
echo Downloading todo tool...
curl -L -o todo.exe %URL%

REM Move the binary to a directory in PATH
echo Installing todo tool to C:\Program Files\...
move todo.exe "C:\Program Files\"

REM Verify installation
where todo.exe >nul 2>nul
if %ERRORLEVEL%==0 (
    echo Installation successful! You can now use the 'todo' command.
) else (
    echo Installation failed. Please check your PATH and try again.
)