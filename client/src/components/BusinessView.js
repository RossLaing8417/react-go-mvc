import React, { useState, useEffect } from 'react';
import { fetchAddresses } from '../services/api';
import AddAddressForm from './AddAddressForm';
import AddressGrid from './AddressGrid';
import EditAddressForm from './EditAddressForm';
import EditBusinessForm from './EditBusinessForm';

const BusinessView = ({ business, onBack, onUpdate }) => {
  const [localBusiness, setLocalBusiness] = useState(business);
  const [addresses, setAddresses] = useState([]);
  const [editingAddress, setEditingAddress] = useState(null);
  const [editingBusiness, setEditingBusiness] = useState(false); // State to track if the business is being edited

  useEffect(() => {
    fetchAddresses(localBusiness.id).then(data => setAddresses(data));
  }, [localBusiness.id]);

  const handleAddAddress = (newAddress) => {
    setAddresses([...addresses, newAddress]);
  };

  const handleEditAddress = (address) => {
    setEditingAddress(address);
  };

  const handleSaveAddress = (updatedAddress) => {
    setAddresses(addresses.map((address) =>
      address.id === updatedAddress.id ? updatedAddress : address
    ));
    setEditingAddress(null);
  };

  const handleDeleteAddress = (addressId) => {
    setAddresses(addresses.filter(address => address.id !== addressId));
  };

  const handleEditBusiness = () => {
    setEditingBusiness(true); // Switch to editing business mode
  };

  const handleSaveBusiness = (updatedBusiness) => {
    setLocalBusiness(updatedBusiness)
    onUpdate(updatedBusiness);
    setEditingBusiness(false); // Switch back to viewing mode
  };

  const handleCancelBusinessEdit = () => {
    setEditingBusiness(false); // Cancel editing
  };

  return (
    <div>
      <button onClick={onBack}>Back to Business List</button>
      <h2>Business Details</h2>
      {editingBusiness ? (
        <EditBusinessForm
          business={localBusiness}
          onSave={handleSaveBusiness}
          onCancel={handleCancelBusinessEdit}
        />
      ) : (
        <>
          <p>Name: {localBusiness.name}</p>
          <p>VAT: {localBusiness.vat_number}</p>
          <p>Registration Number: {localBusiness.registration_number}</p>
          <button onClick={handleEditBusiness}>Edit Business</button>
        </>
      )}

      {editingAddress ? (
        <EditAddressForm
          address={editingAddress}
          onSave={handleSaveAddress}
          onCancel={() => setEditingAddress(null)}
        />
      ) : (
        <div>
          <AddAddressForm businessId={localBusiness.id} onAdd={handleAddAddress} />
          <AddressGrid
            addresses={addresses}
            onEdit={handleEditAddress}
            onDelete={handleDeleteAddress}
          />
        </div>
      )}
    </div>
  );
};

export default BusinessView;
