import { useState } from 'react';

interface LinkCreateFormProps {
  onSubmit: (url: string, title?: string, description?: string, name?: string) => void;
  onCancel: () => void;
}

export const LinkCreateForm = ({ onSubmit, onCancel }: LinkCreateFormProps) => {
  const [url, setUrl] = useState('');
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [name, setName] = useState('');
  const [image, setImage] = useState<File | null>(null);
  const [preview, setPreview] = useState<string | null>(null);
  const [isDragging, setIsDragging] = useState(false);
  const [error, setError] = useState('');

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
    if (file && file.type.startsWith('image/')) {
      handleFileChange(file);
    }
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!url) {
      setError('URL is required');
      return;
    }
    onSubmit(url, title, description, name);
  };

  return (
    <div className="link-create-form">
      <div className="header">
        <h1>OGPer</h1>
        <p className="subtitle">
        Easily add and modify OGP for your URLs
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
            <label htmlFor="image">Image</label>
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
                accept="image/*"
                className="file-upload-input"
                onChange={(e) => handleFileChange(e.target.files?.[0] || null)}
              />
            </div>
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

          {error && <div className="error-message">{error}</div>}

          <div className="footer">
            <div className="footer-content">
              <button type="submit" className="primary">
                Create your link
              </button>
            </div>
          </div>
        </form>
      </div>
    </div>
  );
};