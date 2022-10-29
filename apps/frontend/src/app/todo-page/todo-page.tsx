import React, { useEffect } from 'react';

import { environment } from '../../environments/environment';
import styles from './todo-page.module.scss';

/* eslint-disable-next-line */
export interface TodoPageProps {}

export function TodoPage(props: TodoPageProps) {
  const [tasks, setTasks] = React.useState([]);

  useEffect(() => {
    getData().then((data) => {
      setTasks(data);
    });
  }, []);

  const removeTask = (id: number) => {
    fetch(`${environment.api}/todo/${id}`, {
      method: 'DELETE',
    }).then(() => {
      setTasks(tasks.filter((task: any) => task.ID !== id));
    });
  };

  const handleAddTask = (event: any) => {
    event.preventDefault();

    const task = event.target.elements.task.value;
    event.target.elements.task.value = '';
    fetch(`${environment.api}/todo`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ Title: task }),
    })
      .then((response) => response.json())
      .then((_) => {
        getData().then((data) => {
          setTasks(data);
        });
      });
    console.log(task);
  };

  return (
    <div className={styles['container']}>
      <div className={styles['top-section']}>
        <h1>Welcome to TodoPage!</h1>
      </div>

      <div>
        <form onSubmit={handleAddTask}>
          <label>
            Ajouter une tâche:
            <input
              type="text"
              name="task"
              className={styles['input']}
              id="task"
            />
          </label>
          <button type="submit" className={styles['button']}>
            Ajouter
          </button>
        </form>
      </div>

      {/* TODO: Show current Task */}
      <div>
        <h1>Vos tâches</h1>
        <div className={styles['task-list']}>
          {tasks.map((task: any) => {
            return (
              <div key={task.ID} className={styles['task-item']}>
                {task.Title}
                <button
                  onClick={() => {
                    removeTask(task.ID);
                  }}
                  className={styles['button']}
                >
                  X
                </button>
              </div>
            );
          })}
        </div>
      </div>
    </div>
  );
}

async function getData() {
  const response = await fetch(`${environment.api}/todo`);
  const data = await response.json();
  console.log(data);
  return data;
}

export default TodoPage;
