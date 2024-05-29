// fetchWrapper.js
export function fetchWithAuth(url, options = {}) {
  const token = localStorage.getItem('token');

  const defaultHeaders = {
    'Content-Type': 'application/json',
  };

  const authHeaders = token ? { 'Authorization': `Bearer ${token}` } : {};

  const headers = {
    ...defaultHeaders,
    ...authHeaders,
    ...options.headers,
  };

  const updatedOptions = {
    ...options,
    headers,
  };

  // Using XMLHttpRequest for a synchronous request
  const xhr = new XMLHttpRequest();
  xhr.open(options.method || 'GET', url, false); // `false` makes it synchronous
  for (const [key, value] of Object.entries(headers)) {
    xhr.setRequestHeader(key, value);
  }

  xhr.send(options.body ? JSON.stringify(options.body) : null);

  const response = {
    ok: xhr.status >= 200 && xhr.status < 300,
    status: xhr.status,
    json: () => Promise.resolve(JSON.parse(xhr.responseText)),
    text: () => Promise.resolve(xhr.responseText),
  };

  return response;
}
