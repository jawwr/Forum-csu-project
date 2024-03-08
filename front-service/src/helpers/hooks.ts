import {useCookies} from "react-cookie";

export const useToken = () => {
  const [cookies, setCookies, removeCookies] = useCookies(['FORUM_TOKEN']);

  const { FORUM_TOKEN: token } = cookies;

  const setToken = (value: string) => {
    setCookies('FORUM_TOKEN', value);
  }

  const removeToken = () => {
    removeCookies('FORUM_TOKEN');
  }

  return [
    token,
    setToken,
    removeToken,
  ];
}
