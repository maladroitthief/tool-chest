---
title: Use Case Template Example
description: Example document of the use case template taken from Pragmatic Programmer
author: Ian Weller
date: 09.18.2021
---

# Buy Goods

## Characteristic Information

- **Goal in context:** Buyer issues request directly to our company, expects goods shipped and to be billed.
- **Scope:** Company
- **Level:** Summary
- **Preconditions:** We know the buyer, their address, etc
- **Success end condition:** Buyer has goods, we have money for the goods.
- **Failed end condition:** We have not send the goods, buyer has not sent the money.
- **Primary actor:** Buyer, any agent (or computer) acting for the customer.
- **Trigger:** Purchase request comes in.

## Main Success Scenario

1. Buyer calls in with a purchase request.
2. Company captures the buyer's name, address, requested goods, etc.
3. Company gives buyer information on goods, prices, delivery dates, etc.
   1. Company is out of one of the ordered items: Renegotiate order.
4. Buyer signs for the order.
   1. Buyer pays directly with credit card.
5. Company creates order, ships order to buyer.
6. Company ships invoice to buyer.
7. Buyer pays invoice.
   1. Buyer returns goods.

## Variations

1. Buyer may use phone in, fax in, web interface, electronic exchange.

---

7. Buyer may pay by cash, money order, check, or credit card.

## Related Information

- **Priority:** Top
- **Performance target:** 5 minutes for order, 45 days until paid
- **Frequency:** 200/day
- **Superordinate use case:** Manage customer relationship
- **Subordinate use cases:** Create order, Take payment by credit card, Handle returned goods
- **Channel to primary actor:** May be phone, file, or interactive
- **Secondary actors:** Credit card company, bank, shipping service

## Schedule

- **Due date:** Release 1.0

## Open Issues

- What happens if we have part of the order?
- What happens if credit card is stolen?
