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
  const [error, setError] = useState('');

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
        指定したURLのOGPを簡単に追加・変更できます
        </p>
      </div>

      <div className="form-container">
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="url">URL (required)</label>
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

          <p className="subtitle">
              問い合わせはこちらの<a href="https://twitter.com/ponyo877" target="_blank" rel="noopener noreferrer">X(@ponyo877)</a>まで
          </p>

          {error && <div className="error-message">{error}</div>}

          <div className="form-actions">
            <button type="button" className="secondary" onClick={onCancel}>
              Cancel
            </button>
            <button type="submit" className="primary">
              Start link
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};