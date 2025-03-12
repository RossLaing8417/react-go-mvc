import React, { useEffect, useState } from 'react';
import BusinessGrid from './components/BusinessGrid';
import BusinessView from './components/BusinessView';
import AddBusinessForm from './components/AddBusinessForm';
import { fetchBusinesses } from './services/api';

const App = () => {
  const [businesses, setBusinesses] = useState([]);
  const [selectedBusiness, setSelectedBusiness] = useState(null);

  useEffect(() => {
    fetchBusinesses().then(data => setBusinesses(data));
  }, []);

  const handleAddBusiness = (newBusiness) => {
    setBusinesses([...businesses, newBusiness]);
  };

  const handleSelectBusiness = (business) => {
    setSelectedBusiness(business);
  };

  const handleDeselectBusiness = () => {
    setSelectedBusiness(null);
  };

  const handleUpdateBusiness = (updatedBusiness) => {
    setBusinesses(businesses.map((business) =>
      business.id === updatedBusiness.id ? updatedBusiness : business
    ));
  };

  return (
    <div>
      <h1>Business Management</h1>

      {!selectedBusiness ? (
        <div>
          <AddBusinessForm onAdd={handleAddBusiness} />
          <BusinessGrid businesses={businesses} setBusinesses={setBusinesses} onView={handleSelectBusiness} />
        </div>
      ) : (
        <div>
          <BusinessView business={selectedBusiness} onBack={handleDeselectBusiness} onUpdate={handleUpdateBusiness} />
        </div>
      )}
    </div>
  );
};

export default App;
