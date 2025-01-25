import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { useEffect } from 'react';
import { ulid } from 'ulid';
import { LinkCreatePage } from './pages/LinkCreatePage';
import Sidebar from './components/Sidebar';
import LinksPage from './pages/LinksPage';
import './App.css';

const USER_HASH_KEY = 'user_hash';

function App() {
  useEffect(() => {
    if (!localStorage.getItem(USER_HASH_KEY)) {
      const newUserHash = ulid();
      localStorage.setItem(USER_HASH_KEY, newUserHash);
    }
  }, []);

  const handleSubmit = (destination: string, title?: string, customBackHalf?: string) => {
    console.log('Submitting:', { destination, title, customBackHalf });
    // TODO: Implement actual submission logic
  };

  const handleCancel = () => {
    console.log('Form cancelled');
  };

  return (
    <BrowserRouter>
      <div className="app-layout">
        <Sidebar />
        <div className="main-content">
          <Routes>
            <Route path="/" element={
              <div className="app-container">
                <LinkCreatePage onSubmit={handleSubmit} onCancel={handleCancel} />
              </div>
            } />
            <Route path="/links" element={
              <div className="app-container">
                <LinksPage />
              </div>
            } />
          </Routes>
        </div>
      </div>
    </BrowserRouter>
  );
}

export default App