import { useState } from 'react';
import axios from 'axios';
import { SuccessModal } from '../components/SuccessModal';

const api = axios.create({
  baseURL: import.meta.env.VITE_HOST_NAME || 'http://localhost:8080'
});

interface LinkCreatePageProps {
  onSubmit: (url: string, title?: string, description?: string, name?: string) => void;
  onCancel: () => void;
}

export const LinkCreatePage = ({ onSubmit, onCancel }: LinkCreatePageProps) => {
  const [url, setUrl] = useState('');
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [name, setName] = useState('');
  const [image, setImage] = useState<File | null>(null);
  const [preview, setPreview] = useState<string | null>(null);
  const [isDragging, setIsDragging] = useState(false);
  const [error, setError] = useState('');
  const [showModal, setShowModal] = useState(false);
  const [generatedLink, setGeneratedLink] = useState('');

  const handleFileChange = (file: File | null) => {
    setImage(file);
    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => {
        setPreview(reader.result as string);
      };
      reader.readAsDataURL(file);
    } else {
      setPreview(null);
    }
  };

  const handleDragOver = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    setIsDragging(true);
  };

  const handleDragLeave = () => {
    setIsDragging(false);
  };

  const handleDrop = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    setIsDragging(false);
    const file = e.dataTransfer.files[0];
    if (file && (file.type === 'image/jpeg' || file.type === 'image/png' || file.type === 'image/webp')) {
      handleFileChange(file);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!url) {
      setError('URL is required');
      return;
    }
    if (!image) {
      setError('Image is required');
      return;
    }

    try {
      const userHash = localStorage.getItem('user_hash');
      if (!userHash) {
        setError('User hash not found');
        return;
      }

      const formData = new FormData();
      formData.append('url', url);
      formData.append('title', title);
      formData.append('description', description);
      formData.append('name', name);
      formData.append('user_hash', userHash);
      if (image) {
        formData.append('image', image);
      }

      const response = await api.post('/upload', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      });

      if (response.status === 200) {
        setGeneratedLink(response.data.url);
        setShowModal(true);
        onSubmit(url, title, description, name);
        // Reset form state
        setUrl('');
        setTitle('');
        setDescription('');
        setName('');
        setImage(null);
        setPreview(null);
        setError('');
        
        // Reset file input
        const fileInput = document.getElementById('image') as HTMLInputElement;
        if (fileInput) {
          fileInput.value = '';
        }
      } else {
        setError('Failed to update');
      }
    } catch (error) {
      setError('An error occurred while updating');
      console.error(error);
    }
  };

  return (
    <div className="link-create-form">
      <div className="header">
        <h1>OGPer</h1>
        <p className="subtitle">
          Easily add and modify <span className="tooltip-icon">
            OGP
            <svg xmlns="http://www.w3.org/2000/svg" width="8" height="8" fill="currentColor" viewBox="0 0 16 16">
              <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/>
              <path d="m8.93 6.588-2.29.287-.082.38.45.083c.294.07.352.176.288.469l-.738 3.468c-.194.897.105 1.319.808 1.319.545 0 1.178-.252 1.465-.598l.088-.416c-.2.176-.492.246-.686.246-.275 0-.375-.193-.304-.533L8.93 6.588zM9 4.5a1 1 0 1 1-2 0 1 1 0 0 1 2 0z"/>
            </svg>
          </span> for your URLs
        </p>
      </div>

      <div className="form-container">
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="url">URL</label>
            <input
              id="url"
              type="url"
              value={url}
              onChange={(e) => setUrl(e.target.value)}
              placeholder="https://example.com/my-no-ogp-url"
              autoFocus
              required
            />
          </div>

          <div className="form-group">
            <label htmlFor="image">
              Image
              <span className="image-tooltip">
                <svg xmlns="http://www.w3.org/2000/svg" width="8" height="8" fill="currentColor" viewBox="0 0 16 16">
                  <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/>
                  <path d="m8.93 6.588-2.29.287-.082.38.45.083c.294.07.352.176.288.469l-.738 3.468c-.194.897.105 1.319.808 1.319.545 0 1.178-.252 1.465-.598l.088-.416c-.2.176-.492.246-.686.246-.275 0-.375-.193-.304-.533L8.93 6.588zM9 4.5a1 1 0 1 1-2 0 1 1 0 0 1 2 0z"/>
                </svg>
              </span>
            </label>
            <div
              className={`file-upload-container ${isDragging ? 'dragover' : ''}`}
              onDragOver={handleDragOver}
              onDragLeave={handleDragLeave}
              onDrop={handleDrop}
            >
              <label htmlFor="image" className="file-upload-label">
                {preview ? (
                  <img src={preview} alt="Preview" className="file-preview" />
                ) : (
                  <>
                    <span>Drag & drop or </span>
                    <button
                      type="button"
                      className="secondary"
                      onClick={() => document.getElementById('image')?.click()}
                    >
                      Choose file
                    </button>
                  </>
                )}
              </label>
              <input
                id="image"
                type="file"
                accept="image/jpeg,image/png,image/webp"
                className="file-upload-input"
                onChange={(e) => handleFileChange(e.target.files?.[0] || null)}
              />
            </div>
            {error && <div className="image-error">Image is required</div>}
          </div>

          <div className="form-group">
            <label htmlFor="name">Name (optional)</label>
            <input
              id="name"
              type="text"
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
          </div>

          <div className="form-group">
            <label htmlFor="title">Title (optional)</label>
            <input
              id="title"
              type="text"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
            />
          </div>

          <div className="form-group">
            <label htmlFor="description">Description (optional)</label>
            <textarea
              id="description"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              rows={3}
            />
          </div>

          <div className="contact-info subtitle">
              For inquiries, please contact <a href="https://twitter.com/ponyo877" target="_blank" rel="noopener noreferrer">X(@ponyo877)</a>
          </div>

          <div className="footer">
            <div className="footer-content">
              <button type="submit" className="primary">
                Create your link
              </button>
            </div>
          </div>
          
          {showModal && (
            <SuccessModal
              link={generatedLink}
              onClose={() => setShowModal(false)}
            />
          )}
        </form>
      </div>
    </div>
  );
};