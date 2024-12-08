import { contextBridge, ipcRenderer } from 'electron';

contextBridge.exposeInMainWorld('electronAPI', {
  onSystemStats: (callback: (data: any) => void) => {
    ipcRenderer.on('system-stats', (_, data: any) => {
      if (data.status === 'success') {
        callback(data)
      }
    });
  },
});