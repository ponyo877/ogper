import './OGPCard.css';
import { useState } from 'react';

type OGP = {
  title: string;
  name: string;
  site_url: string;
  ogp_url: string;
  thumbnail: string;
};

interface OGPCardProps {
  ogpData: OGP;
}

const OGPCard = ({ ogpData }: OGPCardProps) => {
  const [copied, setCopied] = useState(false);

  const handleCopy = (text: string) => {
    navigator.clipboard.writeText(text)
      .then(() => {
        setCopied(true);
        setTimeout(() => setCopied(false), 2000);
      })
      .catch((err) => console.error('Failed to copy:', err));
  };

  return (
    <div className="ogp-card">
      <div className="thumbnail-wrapper">
        <img
          src={ogpData.thumbnail}
          alt="サムネイル"
          loading="lazy"
          className="thumbnail"
        />
      </div>
      <div className="content">
        <h3 className="title">{ogpData.title}</h3>
        <p className="name">{ogpData.name}</p>
        <div className="divider">
          <a className="site-name" href={ogpData.ogp_url}>
            {ogpData.ogp_url.replace('https://', '')}
          </a>
          <button
            className="copy-button2"
            onClick={() => handleCopy(ogpData.ogp_url)}
            title="Copy URL"
          >
            {copied ? '✓' : '⎘'}
          </button>
        </div>
        <div className="divider">
          <a className="site-name" href={ogpData.site_url}>
            {ogpData.site_url.replace('https://', '')}
          </a>
          <button
            className="copy-button2"
            onClick={() => handleCopy(ogpData.site_url)}
            title="Copy URL"
          >
            {copied ? '✓' : '⎘'}
          </button>
        </div>
      </div>
    </div>
  );
};

export default OGPCard;