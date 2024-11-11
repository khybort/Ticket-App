import axios from 'axios';
import { Ticket, CreateTicketRequest } from '../types/ticketTypes';

const API_BASE_URL = process.env.REACT_APP_BACKEND_BASE_URL;

const apiClient = axios.create({
  baseURL: API_BASE_URL,
});

export const getTicketDetails = async (ticketId: number) => {
  const response = await apiClient.get(`/api/v1/tickets/${ticketId}`);
  return response.data;
};

export const createTicket = async (
  ticketData: CreateTicketRequest
): Promise<Ticket> => {
  try {
    const response = await apiClient.post(`/api/v1/tickets`, {
      name: ticketData.name,
      desc: ticketData.desc,
      allocation: ticketData.allocation,
    });
    return response.data;
  } catch (error) {
    throw new Error('Error creating ticket. Please try again later.');
  }
};

export const fetchTickets = async (): Promise<Ticket[]> => {
  const response = await apiClient.get(`/api/v1/tickets`);
  return response.data;
};

export const purchaseTicket = async (
  id: number,
  quantity: number,
  userId: string
): Promise<{ status: string }> => {
  const response = await apiClient.post(`/api/v1/tickets/${id}/purchases`, {
    quantity,
    user_id: userId,
  });
  return response.data;
};
