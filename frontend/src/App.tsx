import { useState } from 'react'
import { LinkCreateForm } from './components/LinkCreateForm'
import './App.css'

function App() {
  const handleSubmit = (destination: string, title?: string, customBackHalf?: string) => {
    console.log('Submitting:', { destination, title, customBackHalf });
    // TODO: Implement actual submission logic
  };

  const handleCancel = () => {
    console.log('Form cancelled');
  };

  return (
    <div className="app-container">
      <LinkCreateForm onSubmit={handleSubmit} onCancel={handleCancel} />
    </div>
  )
}

export default App