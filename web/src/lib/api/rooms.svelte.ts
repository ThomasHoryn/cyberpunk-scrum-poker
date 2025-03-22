import { apiRequest } from './client.svelte';

export interface Room {
  id: string;
  name: string;
  createdAt: string;
}

export interface CreateRoomRequest {
  name: string;
}

export interface CreateRoomResponse {
  id: string;
}

export async function createRoom(request: CreateRoomRequest) {
  return apiRequest<CreateRoomResponse>('rooms', {
    method: 'POST',
    body: request
  });
}

export async function getRoom(id: string) {
  return apiRequest<Room>(`rooms/${id}`);
}
