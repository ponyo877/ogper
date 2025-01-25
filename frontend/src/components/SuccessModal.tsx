import { useState } from 'react';

interface SuccessModalProps {
  link: string;
  onClose: () => void;
}

export const SuccessModal = ({ link, onClose }: SuccessModalProps) => {
  const [copied, setCopied] = useState(false);

  const handleCopy = () => {
    navigator.clipboard.writeText(link);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  const shareToTwitter = () => {
    const url = `https://twitter.com/intent/tweet?url=${encodeURIComponent(link)}`;
    window.open(url, '_blank');
  };

  return (
    <div className="modal-overlay">
      <div className="success-modal">
        <h2>Your link is ready!🎉</h2>
        <p className="instruction-text">Copy link below to share it or share on X</p>
        <div className="link-container">
          <a href={link} target="_blank" rel="noopener noreferrer">
            {link.replace('https://', '')}
          </a>
          <button
            type="button"
            onClick={handleCopy}
            className={`copy-button ${copied ? 'copied' : ''}`}
          >
            {copied ? 'Copied!' : 'Copy'}
          </button>
        </div>
        <div className="share-buttons">
          <button type="button" onClick={shareToTwitter} className="twitter-share">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
              <path d="M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z"/>
            </svg>
          </button>
        </div>
        <button onClick={onClose} className="close-icon">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
            <path d="M24 20.188l-8.315-8.209 8.2-8.282-3.697-3.697-8.212 8.318-8.31-8.203-3.666 3.666 8.321 8.24-8.206 8.313 3.666 3.666 8.237-8.318 8.285 8.203z"/>
          </svg>
        </button>
      </div>
    </div>
  );
};