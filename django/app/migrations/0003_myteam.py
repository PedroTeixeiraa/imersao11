# Generated by Django 4.1.4 on 2022-12-10 21:17

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('app', '0002_team'),
    ]

    operations = [
        migrations.CreateModel(
            name='MyTeam',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('players', models.ManyToManyField(to='app.player')),
            ],
        ),
    ]
