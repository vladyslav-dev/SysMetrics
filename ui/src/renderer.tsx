import { createRoot } from 'react-dom/client';
import App from "./App";

interface ElectronAPI {
    onSystemStats: (callback: (data: string) => void) => void;
  }
  
declare global {
  interface Window {
    electronAPI: ElectronAPI;
  }
}

const root = createRoot(document.getElementById('app'));
root.render(<App />);