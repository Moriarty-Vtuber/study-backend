import { createClient } from '@supabase/supabase-js'
import type { SystemConstants, Seat, Menu, WorkNameTrend } from '../types/api'

const supabaseUrl = process.env.NEXT_PUBLIC_SUPABASE_URL!
const supabaseKey = process.env.NEXT_PUBLIC_SUPABASE_ANON_KEY!

export const supabase = createClient(supabaseUrl, supabaseKey)

export const getSystemConstants = async (): Promise<SystemConstants> => {
  const { data, error } = await supabase
    .from('SystemConstants')
    .select('*')
    .single()

  if (error) {
    throw error
  }

  return data
}

export const getSeats = async (): Promise<Seat[]> => {
  const { data, error } = await supabase.from('Seat').select('*')

  if (error) {
    throw error
  }

  return data
}

export const createSeat = async (seat: Seat): Promise<Seat> => {
  const { data, error } = await supabase.from('Seat').insert(seat).single()

  if (error) {
    throw error
  }

  return data
}

export const getMenus = async (): Promise<Menu[]> => {
  const { data, error } = await supabase.from('Menu').select('*')

  if (error) {
    throw error
  }

  return data
}

export const getWorkNameTrends = async (): Promise<WorkNameTrend[]> => {
  const { data, error } = await supabase.from('WorkNameTrend').select('*')

  if (error) {
    throw error
  }

  return data
}
