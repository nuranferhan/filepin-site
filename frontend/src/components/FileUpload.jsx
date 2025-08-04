import React, { useState } from 'react';
import { useDropzone } from 'react-dropzone';
import { uploadFile } from '../services/api';

const FileUpload = ({ onUploadSuccess }) => {
  const [uploading, setUploading] = useState(false);
  const [options, setOptions] = useState({
    expiryHours: 24,
    maxDownloads: 0,
    isOneTime: false,
  });

  const onDrop = async (acceptedFiles) => {
    if (acceptedFiles.length === 0) return;

    setUploading(true);
    try {
      const file = acceptedFiles[0];
      const result = await uploadFile(file, options);
      onUploadSuccess(result);
    } catch (error) {
      alert('Upload failed: ' + error.message);
    } finally {
      setUploading(false);
    }
  };

  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop,
    maxFiles: 1,
  });

  return (
    <div className="upload-container">
      <div className="options-panel">
        <h3>Upload Options</h3>
        <div className="option">
          <label>Expiry Hours:</label>
          <input
            type="number"
            value={options.expiryHours}
            onChange={(e) => setOptions({ ...options, expiryHours: parseInt(e.target.value) })}
            min="1"
            max="168"
          />
        </div>
        <div className="option">
          <label>Max Downloads (0 = unlimited):</label>
          <input
            type="number"
            value={options.maxDownloads}
            onChange={(e) => setOptions({ ...options, maxDownloads: parseInt(e.target.value) })}
            min="0"
          />
        </div>
        <div className="option">
          <label>
            <input
              type="checkbox"
              checked={options.isOneTime}
              onChange={(e) => setOptions({ ...options, isOneTime: e.target.checked })}
            />
            One-time download
          </label>
        </div>
      </div>

      <div {...getRootProps()} className={`dropzone ${isDragActive ? 'active' : ''}`}>
        <input {...getInputProps()} />
        {uploading ? (
          <p>Uploading...</p>
        ) : isDragActive ? (
          <p>Drop the file here...</p>
        ) : (
          <p>Drag & drop a file here, or click to select</p>
        )}
      </div>
    </div>
  );
};

export default FileUpload;