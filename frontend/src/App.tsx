import { BrowserRouter, Route, Routes } from 'react-router-dom';
import type { RouteProps } from 'react-router-dom';
import { LinkCreateForm } from './components/LinkCreateForm';
import Sidebar from './components/Sidebar';
import LinksPage from './pages/LinksPage';
import './App.css';

function App() {
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
                <LinkCreateForm onSubmit={handleSubmit} onCancel={handleCancel} />
              </div>
            } />
            <Route path="/links" element={<LinksPage />} />
          </Routes>
        </div>
      </div>
    </BrowserRouter>
  );
}

export default App