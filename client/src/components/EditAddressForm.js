import React, { useState, useEffect } from 'react';
import { updateAddress } from '../services/api';

const EditAddressForm = ({ address, onCancel, onSave }) => {
  const [streetNumber, setStreetNumber] = useState(address.street_number);
  const [street, setStreet] = useState(address.street);
  const [town, setTown] = useState(address.town);
  const [postCode, setPostCode] = useState(address.post_code);
  const [error, setError] = useState(null);  // State for handling errors

  useEffect(() => {
    setStreetNumber(address.street_number);
    setStreet(address.street);
    setTown(address.town);
    setPostCode(address.post_code);
  }, [address]);

  const handleSubmit = (e) => {
    e.preventDefault();
    const updatedAddress = {
      ...address,
      street_number: streetNumber,
      street,
      town,
      post_code: postCode,
    };

    updateAddress(updatedAddress).then((updated) => {
      onSave(updated); // Update the parent component with the saved address
      setError(null);
    }).catch(err => {
      setError(err || 'Failed to update address');
    });
  };

  return (
    <form onSubmit={handleSubmit}>
      {error && <div style={{ color: 'red' }}><strong>{error.message}</strong></div>}  {/* Display error if it occurs */}
      <div>
        <label htmlFor="streetNumber">Street Number</label>
        <input
          id="streetNumber"
          type="text"
          placeholder="Street Number"
          value={streetNumber}
          onChange={(e) => setStreetNumber(e.target.value)}
          required
        />
      </div>

      <div>
        <label htmlFor="street">Street</label>
        <input
          id="street"
          type="text"
          placeholder="Street"
          value={street}
          onChange={(e) => setStreet(e.target.value)}
          required
        />
      </div>

      <div>
        <label htmlFor="town">Town</label>
        <input
          id="town"
          type="text"
          placeholder="Town"
          value={town}
          onChange={(e) => setTown(e.target.value)}
          required
        />
      </div>

      <div>
        <label htmlFor="postCode">Post Code</label>
        <input
          id="postCode"
          type="text"
          placeholder="Post Code"
          value={postCode}
          onChange={(e) => setPostCode(e.target.value)}
          required
        />
      </div>

      <button type="submit">Save</button>
      <button type="button" onClick={onCancel}>Cancel</button>
    </form>
  );
};

export default EditAddressForm;
