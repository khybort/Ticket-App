export interface Ticket {
  id: number;
  name: string;
  desc: string;
  allocation: number;
}

export interface CreateTicketRequest {
  name: string;
  desc: string;
  allocation: number;
}
