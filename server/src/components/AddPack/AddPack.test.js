import React from 'react';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import AddPack from './AddPack';

describe('<AddPack />', () => {
  test('it should mount', () => {
    render(<AddPack />);
    
    const addPack = screen.getByTestId('AddPack');

    expect(addPack).toBeInTheDocument();
  });
});