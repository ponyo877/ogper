import { useState } from 'react';
import { Link } from 'react-router-dom';
import './Sidebar.css';

const Sidebar = () => {
  const [isCollapsed, setIsCollapsed] = useState(window.innerWidth <= 768);

  const toggleSidebar = () => {
    setIsCollapsed(!isCollapsed);
  };

  return (
    <div className={`sidebar ${isCollapsed ? 'collapsed' : ''}`}>
      <button className="toggle-btn" onClick={toggleSidebar}>
        {isCollapsed ? '▶' : '◀'}
      </button>
      <nav>
        <ul>
          <li>
            <Link to="/">
              <span role="img" aria-label="Create">✏️</span>
              {!isCollapsed && ' Create'}
            </Link>
          </li>
          <li>
            <Link to="/links">
              <span role="img" aria-label="Links">🔗</span>
              {!isCollapsed && ' Links'}
            </Link>
          </li>
        </ul>
      </nav>
    </div>
  );
};

export default Sidebar;