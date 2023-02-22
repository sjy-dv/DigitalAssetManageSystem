@echo off

rem Run the wmic command to get disk usage information
for /f "usebackq skip=1 tokens=1,2" %%i in (`wmic logicaldisk where "deviceid='C:'" get freespace^,size`) do (
    set free_space=%%i
    set total_space=%%j
)

set /a free_space_gb=%free_space%/1024/1024/1024

if %free_space_gb% gtr 5 (
    echo OK
) else (
    echo Low disk space: %free_space_gb% GB
)
