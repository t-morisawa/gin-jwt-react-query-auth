import { storage } from './utils';

interface AuthResponse {
  user: User;
  jwt: string;
}

export interface User {
  id: string;
  email: string;
  name?: string;
}

const API_BASE_URL = process.env.REACT_APP_API_URL

export async function handleApiResponse(response) {
  const data = await response.json();

  if (response.ok) {
    return data;
  } else {
    return Promise.reject(data);
  }
}

export async function getUserProfile() {
  return await fetch(API_BASE_URL + '/auth/me', {
    headers: {
      Authorization: storage.getToken(),
    },
  }).then(handleApiResponse);
}

export async function loginWithEmailAndPassword(data): Promise<AuthResponse> {
  return window
    .fetch(API_BASE_URL + '/auth/login', {
      method: 'POST',
      body: JSON.stringify(data),
    })
    .then(handleApiResponse);
}

export async function registerWithEmailAndPassword(
  data
): Promise<AuthResponse> {
  return window
    .fetch(API_BASE_URL + '/auth/register', {
      method: 'POST',
      body: JSON.stringify(data),
    })
    .then(handleApiResponse);
}
