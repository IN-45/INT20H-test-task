import axios, { AxiosRequestConfig, AxiosResponse, Method } from 'axios';

const apiAxios = axios.create({
  baseURL: 'http://localhost:8000',
});

export async function requestApi<T = never, D = never>(
  method: Method,
  url: string,
  config?: AxiosRequestConfig<D>,
): Promise<AxiosResponse<T>> {
  return await apiAxios<T, AxiosResponse<T, D>, D>({ ...config, method, url });
}
