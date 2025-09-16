import asyncio
import json
import sys
import win32pipe
import win32file

PIPE_NAME = r'\\.\pipe\AppPulsePipe'

async def read_pipe():
    print("Connecting to pipe...")
    try:
        handle = win32file.CreateFile(
            PIPE_NAME,
            win32file.GENERIC_READ,
            0,
            None,
            win32file.OPEN_EXISTING,
            0,
            None
        )
    except Exception as e:
        if e.winerror == 2:
            print("Pipe not found. Is the C# application running?")
        else:
            print(f"Error opening pipe: {e}")
        return

    print("Connected to pipe. Waiting for messages...")
    
    try:
        while True:
            try:
                # Read from the pipe
                resp = win32file.ReadFile(handle, 4096)
                
                # resp is a tuple (hr, data)
                if resp[0] == 0: # ERROR_SUCCESS
                    raw = resp[1]
                    decoded = None
                    used_enc = None
                    # Try common encodings: utf-8, cp950 (Big5), big5, then latin-1 as last resort
                    for enc in ('utf-8', 'cp950', 'big5', 'latin-1'):
                        try:
                            decoded = raw.decode(enc).strip()
                            used_enc = enc
                            break
                        except (UnicodeDecodeError, LookupError):
                            continue
                    if decoded is None:
                        # Fallback: print hex dump to help debugging
                        hexdump = ' '.join(f'{b:02x}' for b in raw)
                        print(f"Unable to decode message bytes; hex: {hexdump}")
                        decoded = raw.decode('latin-1', errors='ignore').strip()
                        used_enc = 'latin-1 (fallback, ignored errors)'
                    if decoded:
                        lines = decoded.split('\n')
                        for line in lines:
                            if line:
                                try:
                                    data = json.loads(line)
                                    print(json.dumps(data, indent=2, ensure_ascii=False))
                                except json.JSONDecodeError:
                                    print(f"Received non-JSON message ({used_enc}): {line}")
                else:
                    # Handle other potential errors if needed
                    await asyncio.sleep(0.1)

            except win32file.error as e:
                # ERROR_BROKEN_PIPE (109) means the server disconnected.
                if e.winerror == 109:
                    print("Pipe closed by server. Reconnecting...")
                    win32file.CloseHandle(handle)
                    while True:
                        try:
                            handle = win32file.CreateFile(
                                PIPE_NAME,
                                win32file.GENERIC_READ,
                                0,
                                None,
                                win32file.OPEN_EXISTING,
                                0,
                                None
                            )
                            print("Reconnected.")
                            break
                        except Exception:
                            await asyncio.sleep(1) # Wait before retrying
                else:
                    print(f"Error reading from pipe: {e}")
                    break
            
            await asyncio.sleep(0.01) # Small delay to prevent busy-waiting

    except KeyboardInterrupt:
        print("Stopping client.")
    finally:
        win32file.CloseHandle(handle)

if __name__ == "__main__":
    if sys.platform != "win32":
        print("This script is designed to run on Windows.")
        sys.exit(1)
        
    try:
        asyncio.run(read_pipe())
    except KeyboardInterrupt:
        print("\nExiting.")
