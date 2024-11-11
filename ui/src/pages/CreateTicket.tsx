import React, { useState } from 'react';
import axios from 'axios';
import { toast } from 'react-toastify';
import { useNavigate } from 'react-router-dom';
import Navbar from '../components/Navbar/Navbar';
import { createTicket } from '../services/ticketService';
import { CreateTicketRequest } from '../types/ticketTypes';
import 'react-toastify/dist/ReactToastify.css';

const CreateTicket: React.FC = () => {
  const [ticketName, setTicketName] = useState('');
  const [ticketDesc, setTicketDesc] = useState('');
  const [allocation, setAllocation] = useState<number>(0);
  const navigate = useNavigate();

  const handleCreateTicket = async () => {
    if (!ticketName || !ticketDesc || allocation <= 0) {
      toast.error('Please fill in all fields correctly.');
      return;
    }

    try {
      const createTicketRequest: CreateTicketRequest = {
        name: ticketName,
        desc: ticketDesc,
        allocation,
      };
      await createTicket(createTicketRequest);
      toast.success('Ticket Created Successfully!');

      navigate('/ticket-list');
    } catch (error: unknown) {
      if (axios.isAxiosError(error)) {
        toast.error(error.response?.data?.message || 'Error creating ticket.');
      } else {
        toast.error('An unexpected error occurred.');
      }
    }
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />
      <div className="container mx-auto p-8">
        <h1 className="text-3xl font-bold text-center mb-6">
          Create New Ticket
        </h1>
        <div className="space-y-4">
          <input
            type="text"
            value={ticketName}
            onChange={(e) => setTicketName(e.target.value)}
            placeholder="Ticket Name"
            className="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          <textarea
            value={ticketDesc}
            onChange={(e) => setTicketDesc(e.target.value)}
            placeholder="Ticket Description"
            className="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          <input
            type="number"
            value={allocation}
            onChange={(e) => setAllocation(Number(e.target.value))}
            placeholder="Ticket Allocation"
            className="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          <button
            type="button"
            onClick={handleCreateTicket}
            className="w-full bg-green-500 text-white px-6 py-2 rounded-lg hover:bg-green-600"
          >
            Create Ticket
          </button>
        </div>
      </div>
    </div>
  );
};

export default CreateTicket;
