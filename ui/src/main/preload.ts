import { contextBridge, ipcRenderer } from 'electron';

contextBridge.exposeInMainWorld('electronAPI', {
  onSystemStats: (callback: (data: string) => void) => {
    ipcRenderer.on('system-stats', (_, data: string) => {
      console.log('data', data)
      callback(data)
    });
  },
});