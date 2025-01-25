import { useEffect, useState } from 'react'
import OGPCard from '../components/OGPCard';
import axios from 'axios';
import './LinksPage.css';

const api = axios.create({
  baseURL: import.meta.env.VITE_HOST_NAME || 'http://localhost:8080'
});

interface Link {
  ogper_url: string;
  title?: string;
  description?: string;
  name?: string;
  site_url: string;
  image_url: string;
  published_at: string;
}

const LinksPage = () => {
  const [links, setLinks] = useState<Link[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchLinks = async () => {
      try {
        const userHash = localStorage.getItem('user_hash');
        if (!userHash) {
          throw new Error('User hash not found');
        }

        const response = await api.get(`/links?user_hash=${userHash}`);
        setLinks(response.data);
      } catch (err) {
        setError('Failed to fetch links');
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    fetchLinks();
  }, []);

  const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    if (isNaN(date.getTime())) {
      return 'Invalid Date';
    }
    return date.toLocaleString('ja-JP', {
      timeZone: 'Asia/Tokyo',
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit'
    });
  };

  if (loading) {
    return <div className="link-create-form"></div>
  }

  if (error) {
    return <div className="link-create-form">{error}</div>;
  }

  return (
    <div className="link-create-form">
      <div className="header">
        <h1>Links</h1>
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
        {links.length === 0 ? (
          <p className="no-links">No links found</p>
        ) : (
          <>
            {links.map((link) => {
              const ogpData = {
                title: link.title || 'No Title',
                name: link.name || 'No Name',
                site_url: link.site_url,
                ogp_url: link.ogper_url,
                thumbnail: link.image_url
              };
              return (
                <OGPCard key={link.ogper_url} ogpData={ogpData} />
              );
            })}
            </>
        )}
    </div>
  );
};

export default LinksPage;