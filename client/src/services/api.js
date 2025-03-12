import { API_BASE_URL, BUSINESS_API, ADDRESS_API } from './config';

const handleResponse = response => {
  if (!response.ok) {
    return response.json().then(errorData => {
      throw new Error(errorData.error || 'An unknown error occurred');
    }).catch(error => {
      throw error;
    });
  }
  if (response.status === 204) {
    return Promise.resolve();
  }
  return response.json();
};

export const fetchBusinesses = async () => {
  const response = await fetch(`${API_BASE_URL}${BUSINESS_API}`);
  return handleResponse(response);
};

export const addBusiness = async (businessData) => {
  const response = await fetch(`${API_BASE_URL}${BUSINESS_API}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(businessData),
  });
  return handleResponse(response);
};

export const updateBusiness = async (updatedBusiness) => {
  const response = await fetch(`${API_BASE_URL}${BUSINESS_API}/${updatedBusiness.id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(updatedBusiness),
  });
  return handleResponse(response);
};

export const deleteBusiness = async (businessId) => {
  const response = await fetch(`${API_BASE_URL}${BUSINESS_API}/${businessId}`, {
    method: 'DELETE',
  });
  return handleResponse(response);
};

export const fetchAddresses = async (businessId) => {
  const response = await fetch(`${API_BASE_URL}${ADDRESS_API}?business_id=${businessId}`);
  return handleResponse(response);
};

export const addAddress = async (addressData) => {
  const response = await fetch(`${API_BASE_URL}${ADDRESS_API}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(addressData),
  });
  return handleResponse(response);
};

export const updateAddress = async (updatedAddress) => {
  const response = await fetch(`${API_BASE_URL}${ADDRESS_API}/${updatedAddress.id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(updatedAddress),
  });
  return handleResponse(response);
};

export const deleteAddress = async (addressId) => {
  const response = await fetch(`${API_BASE_URL}${ADDRESS_API}/${addressId}`, {
    method: 'DELETE',
  });
  return handleResponse(response);
};
