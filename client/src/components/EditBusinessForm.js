import React, { useState, useEffect } from 'react';
import { updateBusiness } from '../services/api';

const EditBusinessForm = ({ business, onSave, onCancel }) => {
  const [name, setName] = useState(business.name);
  const [vatNumber, setVatNumber] = useState(business.vat_number);
  const [registrationNumber, setRegistrationNumber] = useState(business.registration_number);
  const [error, setError] = useState(null);  // State for handling errors

  useEffect(() => {
    setName(business.name);
    setVatNumber(business.vat_number);
    setRegistrationNumber(business.registration_number);
  }, [business]);

  const handleSubmit = (e) => {
    e.preventDefault();

    const updatedBusiness = {
      ...business,
      name,
      vat_number: Number(vatNumber),
      registration_number: registrationNumber,
    };

    updateBusiness(updatedBusiness).then((updated) => {
      onSave(updated); // Call onSave to update the parent component with the updated business
      setError(null);
    }).catch(err => {
      setError(err || 'Failed to update business');
    });
  };

  return (
    <form onSubmit={handleSubmit}>
      {error && <div style={{ color: 'red' }}><strong>{error.message}</strong></div>}  {/* Display error if it occurs */}
      <div>
        <label htmlFor="name">Business Name</label>
        <input
          id="name"
          type="text"
          placeholder="Business Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />
      </div>

      <div>
        <label htmlFor="vatNumber">VAT Number</label>
        <input
          id="vatNumber"
          type="number"
          placeholder="VAT Number"
          value={vatNumber}
          onChange={(e) => setVatNumber(e.target.value)}
          required
        />
      </div>

      <div>
        <label htmlFor="registrationNumber">Registration Number</label>
        <input
          id="registrationNumber"
          type="text"
          placeholder="Registration Number"
          value={registrationNumber}
          onChange={(e) => setRegistrationNumber(e.target.value)}
          required
        />
      </div>

      <button type="submit">Save</button>
      <button type="button" onClick={onCancel}>Cancel</button>
    </form>
  );
};

export default EditBusinessForm;
