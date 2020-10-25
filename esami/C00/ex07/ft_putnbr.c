/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_putnbr.c                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: abreglia <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/10/25 11:04:18 by abreglia          #+#    #+#             */
/*   Updated: 2020/10/25 18:17:46 by abreglia         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include <unistd.h>
void	write_num();

void	ft_putnbr(int nb)
{
	char arr[100];

	int o;

	int size;

	size = 0;

	while(nb != 0)
	{
		if (nb < 0)
		{
			write(1, "-", 1);
			nb = nb * -1;
		}
		o = nb % 10;
		arr[size] = o + '0';
		size++;
		nb = nb / 10;
	}
		write_num(arr, size);
}

void	write_num(char arr[], int size)
{
	char a;

	while(size >= 0)
	{
		a = arr[size];
		write(1, &a, 1);
		size--;
	}
}
