import { Console, Command } from 'nestjs-console';
import { getConnection } from 'typeorm';
import * as chalk from 'chalk';

import fakeBanks from './fake-bank/bank-001';

@Console()
export class FakeCommand {
  @Command({
    command: 'fake',
    description: 'Seed data in database',
  })
  async command() {
    await this.runMigrations();
    // const fakes = (await import(`./fake-bank/bank-${process.env.BANK_CODE}`))
    //   .default;
    for (const bank of fakeBanks) {
      await this.createInDatabase(bank.model, bank.fields);
    }

    console.log(chalk.green('Data generated'));
  }

  async runMigrations() {
    const conn = getConnection('default');
    for (const _ of conn.migrations.reverse()) {
      await conn.undoLastMigration();
    }
  }

  async createInDatabase(model: any, data: any) {
    const repository = this.getRepository(model);
    const obj = repository.create(data);
    await repository.save(obj);
  }

  getRepository(model: any) {
    const conn = getConnection('default');
    return conn.getRepository(model);
  }
}
