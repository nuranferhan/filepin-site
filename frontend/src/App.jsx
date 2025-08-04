import React, { useState } from 'react';
import FileUpload from './components/FileUpload';
import FileList from './components/FileList';
import FileDownload from './components/FileDownload';
import './App.css';

function App() {
  const [activeTab, setActiveTab] = useState('upload');
  const [refreshTrigger, setRefreshTrigger] = useState(0);

  const handleUploadSuccess = (result) => {
    alert(`File uploaded successfully! Share link: ${window.location.origin}${result.share_link}`);
    setRefreshTrigger(prev => prev + 1);
  };

  return (
    <div className="App">
      <header className="app-header">
        <h1>SecureShare</h1>
        <p>Secure File Sharing with End-to-End Encryption</p>
      </header>

      <nav className="tab-nav">
        <button 
          className={activeTab === 'upload' ? 'active' : ''}
          onClick={() => setActiveTab('upload')}
        >
          Upload
        </button>
        <button 
          className={activeTab === 'download' ? 'active' : ''}
          onClick={() => setActiveTab('download')}
        >
          Download
        </button>
        <button 
          className={activeTab === 'files' ? 'active' : ''}
          onClick={() => setActiveTab('files')}
        >
          My Files
        </button>
      </nav>

      <main className="main-content">
        {activeTab === 'upload' && (
          <FileUpload onUploadSuccess={handleUploadSuccess} />
        )}
        {activeTab === 'download' && (
          <FileDownload />
        )}
        {activeTab === 'files' && (
          <FileList refreshTrigger={refreshTrigger} />
        )}
      </main>

      <footer className="app-footer">
        <p>© 2024 SecureShare - Your files are encrypted and secure</p>
      </footer>
    </div>
  );
}

export default App;