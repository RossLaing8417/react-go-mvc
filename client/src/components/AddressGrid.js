import React, { useState } from 'react';
import { deleteAddress } from '../services/api';

const AddressGrid = ({ addresses, onEdit, onDelete }) => {
  const [error, setError] = useState(null);  // State for handling errors

  const handleDelete = (addressId) => {
    deleteAddress(addressId).then(() => {
      onDelete(addressId); // Call onDelete to update the parent component's address list
      setError(null);
    }).catch(err => {
      setError(err || 'Failed to delete address');
    });
  };

  return (
    <div>
      <h3>Addresses</h3>
      {error && <div style={{ color: 'red' }}><strong>{error.message}</strong></div>}  {/* Display error if it occurs */}
      <table className="address-table">
        <thead>
          <tr>
            <th>Street Number</th>
            <th>Street</th>
            <th>Town</th>
            <th>Post Code</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {addresses.map((address) => (
            <tr key={address.id}>
              <td>{address.street_number}</td>
              <td>{address.street}</td>
              <td>{address.town}</td>
              <td>{address.post_code}</td>
              <td>
                <button onClick={() => onEdit(address)}>Edit</button>
                <button onClick={() => handleDelete(address.id)}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default AddressGrid;
