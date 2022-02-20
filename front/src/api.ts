import { storage } from './utils';

interface AuthResponse {
  user: User;
  jwt: string;
}

export interface User {
  email: string;
  name?: string;
}

export async function handleApiResponse(response) {
  const data = await response.json();

  if (response.ok) {
    return data;
  } else {
    return Promise.reject(data);
  }
}

export async function getUserProfile() {
  return await fetch('/auth/me', {
    headers: {
      Authorization: 'Bearer ' + storage.getToken(),
      'Content-Type': 'application/json',
    },
  }).then(handleApiResponse);
}

export async function loginWithEmailAndPassword(data): Promise<AuthResponse> {
  return window
    .fetch('/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },  
      body: JSON.stringify(data),
    })
    .then(handleApiResponse);
}

export async function registerWithEmailAndPassword(
  data
): Promise<AuthResponse> {
  return window
    .fetch('/auth/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },  
      body: JSON.stringify(data),
    })
    .then(handleApiResponse);
}
