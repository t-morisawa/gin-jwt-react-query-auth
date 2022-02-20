export const storage = {
  getToken: () => JSON.parse(window.localStorage.getItem('token')),
  setToken: token =>
    window.localStorage.setItem('token', JSON.stringify(token)),
  clearToken: () => window.localStorage.removeItem('token'),
};

export const get_token_from_res = (body: Object) => {
    return body['token'];
}
