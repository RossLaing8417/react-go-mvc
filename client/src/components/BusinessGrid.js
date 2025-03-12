import React, { useState } from 'react';
import { deleteBusiness } from '../services/api';

const BusinessGrid = ({ businesses, setBusinesses, onView }) => {
  const [error, setError] = useState(null);  // State for handling errors

  const handleDelete = (id) => {
    deleteBusiness(id).then(() => {
      setBusinesses(businesses.filter(business => business.id !== id));
      setError(null);
    }).catch(err => {
      setError(err || 'Failed to delete business');
    });
  };

  return (
    <div>
      <h1>Businesses</h1>
      {error && <div style={{ color: 'red' }}><strong>{error.message}</strong></div>}  {/* Display error if it occurs */}
      <table className="business-table">
        <thead>
          <tr>
            <th>Name</th>
            <th>VAT Number</th>
            <th>Registration Number</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {businesses.map((business) => (
            <tr key={business.id}>
              <td>{business.name}</td>
              <td>{business.vat_number}</td>
              <td>{business.registration_number}</td>
              <td>
                <button onClick={() => onView(business)}>View</button>
                <button onClick={() => handleDelete(business.id)}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default BusinessGrid;
