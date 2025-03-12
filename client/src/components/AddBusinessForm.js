import React, { useState } from 'react';
import { addBusiness } from '../services/api';

const AddBusinessForm = ({ onAdd }) => {
  const [name, setName] = useState('');
  const [vatNumber, setVatNumber] = useState('');
  const [registrationNumber, setRegistrationNumber] = useState('');
  const [error, setError] = useState(null);  // State for handling errors

  const handleSubmit = (e) => {
    e.preventDefault();
    if (isNaN(vatNumber)) {
      alert("VAT Number must be numeric");
      return;
    }
    const newBusiness = { name, vat_number: Number(vatNumber), registration_number: registrationNumber };
    addBusiness(newBusiness).then((business) => {
      onAdd(business);
      setError(null);
    }).catch(err => {
      setError(err || 'Failed to add business');
    });
  };

  return (
    <form onSubmit={handleSubmit}>
      {error && <div style={{ color: 'red' }}><strong>{error.message}</strong></div>}  {/* Display error if it occurs */}
      <input
        type="text"
        placeholder="Business Name"
        value={name}
        onChange={(e) => setName(e.target.value)}
      />
      <input
        type="number"
        placeholder="VAT Number"
        value={vatNumber}
        onChange={(e) => setVatNumber(e.target.value)}
        required
      />
      <input
        type="text"
        placeholder="Registration Number"
        value={registrationNumber}
        onChange={(e) => setRegistrationNumber(e.target.value)}
      />
      <button type="submit">Add Business</button>
    </form>
  );
};

export default AddBusinessForm;
