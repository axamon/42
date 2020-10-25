/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_putnbr.c                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: abreglia <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/10/25 11:04:18 by abreglia          #+#    #+#             */
/*   Updated: 2020/10/25 17:15:30 by abreglia         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include <unistd.h>

void	ft_putnbr(int nb)
{
	char arr[100];

	char a;

	int o;

	int size;

	size = 0;

	while(nb != 0)
	{
		o = nb % 10;
		arr[size] = o + '0';
		size++;
		nb = nb / 10;
	}
	while(size >=0)
	{	a = arr[size];
		write(1, &a, 1);
		size--;
	}
}
