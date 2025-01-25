import './OGPCard.css';

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
        <a className="site-name" href={ogpData.ogp_url}>
          {ogpData.ogp_url.replace('https://', '')}
        </a>
        <a className="site-name" href={ogpData.site_url}>
          {ogpData.site_url.replace('https://', '')}
        </a>
      </div>
    </div>
  );
};

export default OGPCard;