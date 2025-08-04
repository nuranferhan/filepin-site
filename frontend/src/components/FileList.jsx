import React, { useState, useEffect } from 'react';
import { listFiles, deleteFile } from '../services/api';

const FileList = ({ refreshTrigger }) => {
  const [files, setFiles] = useState([]);
  const [loading, setLoading] = useState(false);

  const loadFiles = async () => {
    setLoading(true);
    try {
      const response = await listFiles();
      setFiles(response.files || []);
    } catch (error) {
      console.error('Failed to load files:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async (fileId) => {
    if (!window.confirm('Are you sure you want to delete this file?')) return;

    try {
      await deleteFile(fileId);
      loadFiles();
    } catch (error) {
      alert('Delete failed: ' + error.message);
    }
  };

  const copyShareLink = (fileId) => {
    const link = `${window.location.origin}/download/${fileId}`;
    navigator.clipboard.writeText(link);
    alert('Share link copied to clipboard!');
  };

  useEffect(() => {
    loadFiles();
  }, [refreshTrigger]);

  if (loading) return <div>Loading files...</div>;

  return (
    <div className="file-list">
      <h3>Uploaded Files</h3>
      {files.length === 0 ? (
        <p>No files uploaded yet.</p>
      ) : (
        <div className="files-grid">
          {files.map((file) => (
            <div key={file.id} className="file-item">
              <div className="file-info">
                <h4>{file.original_name}</h4>
                <p>Size: {(file.size / 1024).toFixed(2)} KB</p>
                <p>Expires: {new Date(file.expires_at).toLocaleString()}</p>
                <p>Downloads: {file.download_count}/{file.max_downloads || '∞'}</p>
              </div>
              <div className="file-actions">
                <button onClick={() => copyShareLink(file.id)}>Copy Link</button>
                <button onClick={() => handleDelete(file.id)} className="delete-btn">
                  Delete
                </button>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default FileList;