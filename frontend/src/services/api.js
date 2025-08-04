import axios from 'axios';

const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080/api/v1';

const api = axios.create({
  baseURL: API_BASE_URL,
});

export const uploadFile = async (file, options = {}) => {
  const formData = new FormData();
  formData.append('file', file);
  formData.append('expiry_hours', options.expiryHours || 24);
  formData.append('max_downloads', options.maxDownloads || 0);
  formData.append('is_one_time', options.isOneTime || false);

  const response = await api.post('/upload', formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
  return response.data;
};

export const downloadFile = async (fileId) => {
  const response = await api.get(`/download/${fileId}`, {
    responseType: 'blob',
  });
  return response;
};

export const listFiles = async () => {
  const response = await api.get('/files');
  return response.data;
};

export const deleteFile = async (fileId) => {
  const response = await api.delete(`/files/${fileId}`);
  return response.data;
};