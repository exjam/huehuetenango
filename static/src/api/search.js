import axios from 'axios';
import { apiAddress } from './config';

export default {
  async perform(term) {
    return await axios.post(`${apiAddress}/api/search`, { term });
  },
};
