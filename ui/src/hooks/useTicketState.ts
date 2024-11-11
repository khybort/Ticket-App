import { useState, useEffect } from 'react';
import {
  fetchTickets,
  purchaseTicket,
  getTicketDetails,
} from '../services/ticketService';
import { Ticket } from '../types/ticketTypes';

const useTicketState = () => {
  const [tickets, setTickets] = useState<Ticket[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [userId, setUserId] = useState<string>('');
  const [quantities, setQuantities] = useState<{ [key: number]: number }>({});
  const [selectedTicket, setSelectedTicket] = useState<Ticket | null>(null);
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false);

  useEffect(() => {
    fetchTickets()
      .then(setTickets)
      .catch(() => setError('Failed to load tickets.'));
  }, []);

  const handlePurchase = async (ticketId: number, quantity: number) => {
    if (!userId) {
      setError('Please provide a valid user ID.');
      return;
    }

    setLoading(true);
    setError(null);

    try {
      await purchaseTicket(ticketId, quantity, userId);
      setTickets((prevTickets) =>
        prevTickets.map((ticket) =>
          ticket.id === ticketId && ticket.allocation >= quantity
            ? { ...ticket, allocation: ticket.allocation - quantity }
            : ticket
        )
      );
      setQuantities((prevQuantities) => ({ ...prevQuantities, [ticketId]: 1 }));
    } catch {
      setError('Failed to purchase the ticket. Please try again later.');
    } finally {
      setLoading(false);
    }
  };

  const handleQuantityChange = (ticketId: number, value: number) => {
    setQuantities((prevQuantities) => ({
      ...prevQuantities,
      [ticketId]: value,
    }));
  };

  const handleGetTicket = async (ticketId: number) => {
    setSelectedTicket(null);
    try {
      const ticketDetails = await getTicketDetails(ticketId);
      setSelectedTicket(ticketDetails);
      setIsModalOpen(true);
    } catch {
      setError('Failed to fetch ticket details.');
    }
  };

  const closeModal = () => {
    setIsModalOpen(false);
    setSelectedTicket(null);
  };

  return {
    tickets,
    loading,
    error,
    userId,
    quantities,
    selectedTicket,
    isModalOpen,
    handlePurchase,
    handleQuantityChange,
    handleGetTicket,
    closeModal,
    setUserId,
  };
};

export default useTicketState;
