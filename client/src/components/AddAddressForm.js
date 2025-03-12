import React, { useState } from 'react';
import { addAddress } from '../services/api';

const AddAddressForm = ({ businessId, onAdd }) => {
  const [streetNumber, setStreetNumber] = useState('');
  const [street, setStreet] = useState('');
  const [town, setTown] = useState('');
  const [postCode, setPostCode] = useState('');
  const [error, setError] = useState(null);  // State for handling errors

  const handleSubmit = (e) => {
    e.preventDefault();
    const newAddress = { business_id: businessId, street_number: streetNumber, street, town, post_code: postCode };
    addAddress(newAddress).then((address) => {
      onAdd(address);
      setError(null);
    }).catch(err => {
      setError(err || 'Failed to add address');
    });
  };

  return (
    <form onSubmit={handleSubmit}>
      {error && <div style={{ color: 'red' }}><strong>{error.message}</strong></div>}  {/* Display error if it occurs */}
      <input
        type="text"
        placeholder="Street Number"
        value={streetNumber}
        onChange={(e) => setStreetNumber(e.target.value)}
      />
      <input
        type="text"
        placeholder="Street"
        value={street}
        onChange={(e) => setStreet(e.target.value)}
      />
      <input
        type="text"
        placeholder="Town"
        value={town}
        onChange={(e) => setTown(e.target.value)}
      />
      <input
        type="text"
        placeholder="Post Code"
        value={postCode}
        onChange={(e) => setPostCode(e.target.value)}
      />
      <button type="submit">Add Address</button>
    </form>
  );
};

export default AddAddressForm;
