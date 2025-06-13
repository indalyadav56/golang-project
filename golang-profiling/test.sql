select * from users left join departments d on d.id = user.id


SELECT MAX(salary)
FROM employees
WHERE salary < (SELECT MAX(salary) FROM employees);



select Max(salary)
from employees
WHERE salery < (select Max(salery) from employees)



-- select employess e where e.salery = 