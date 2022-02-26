import { initReactQueryAuth } from 'react-query-auth';
import {
  getUserProfile,
  registerWithEmailAndPassword,
  loginWithEmailAndPassword,
  User,
} from '../api';
import { storage } from '../utils';
import Loader from '../components/Loader';
import ErrorComponent from '../components/ErrorComponent';

export type LoginCredentials = {
  email: string;
  password: string;
};

export type RegisterCredentials = {
  email: string;
  name: string;
  password: string;
};

async function handleUserResponse(data) {
  const { token } = data;
  storage.setToken(token);
  return;
}

async function loadUser() {
  let user = null;

  if (storage.getToken()) {
    const data = await getUserProfile();
    user = data;
  }
  return user;
}

async function loginFn(data: LoginCredentials) {
  const response = await loginWithEmailAndPassword(data);
  handleUserResponse(response);
  const user = await loadUser();
  return user;
}

async function registerFn(data: RegisterCredentials) {
  const response = await registerWithEmailAndPassword(data);
  handleUserResponse(response);
  const user = await loadUser();
  return user;
}

async function logoutFn() {
  await storage.clearToken();
}

const authConfig = {
  loadUser,
  loginFn,
  registerFn,
  logoutFn,
  "LoaderComponent": Loader,
  ErrorComponent,
};

const { AuthProvider, AuthConsumer, useAuth } = initReactQueryAuth<
  User,
  any,
  LoginCredentials,
  RegisterCredentials
>(authConfig);

export { AuthProvider, AuthConsumer, useAuth };
