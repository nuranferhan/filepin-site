import React, { useState } from 'react';
import { downloadFile } from '../services/api';

const FileDownload = () => {
  const [fileId, setFileId] = useState('');
  const [downloading, setDownloading] = useState(false);

  const handleDownload = async () => {
    if (!fileId.trim()) return;

    setDownloading(true);
    try {
      const response = await downloadFile(fileId);
      
      const blob = new Blob([response.data]);
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      
      const contentDisposition = response.headers['content-disposition'];
      const filename = contentDisposition 
        ? contentDisposition.split('filename=')[1].replace(/"/g, '')
        : 'downloaded-file';
      
      a.download = filename;
      document.body.appendChild(a);
      a.click();
      window.URL.revokeObjectURL(url);
      document.body.removeChild(a);
    } catch (error) {
      alert('Download failed: ' + error.message);
    } finally {
      setDownloading(false);
    }
  };

  return (
    <div className="download-container">
      <h3>Download File</h3>
      <div className="download-form">
        <input
          type="text"
          placeholder="Enter file ID or paste share link"
          value={fileId}
          onChange={(e) => {
            let value = e.target.value;
            if (value.includes('/download/')) {
              value = value.split('/download/')[1];
            }
            setFileId(value);
          }}
        />
        <button 
          onClick={handleDownload} 
          disabled={!fileId.trim() || downloading}
        >
          {downloading ? 'Downloading...' : 'Download'}
        </button>
      </div>
    </div>
  );
};

export default FileDownload;