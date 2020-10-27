/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_print_combn.c                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: abreglia <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/10/27 11:23:32 by abreglia          #+#    #+#             */
/*   Updated: 2020/10/27 11:29:48 by abreglia         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include <unistd.h>

void	ft_print_combn(int n)
{
	char c[n];
	
	int i;

	i = 0;

	while(i < n)
	{
		c[i] = i;
		i++;
	}
		write(1, &c, 3);
}
