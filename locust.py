from locust import HttpUser, TaskSet, task,between


class QuickstartUser(HttpUser):
    wait_time = between(1, 2)

    @task
    def GetSoldiersByCvc(self):
       self.client.get("/GetSoldiersByCvc?cvc=1000")
    @task
    def GetCombatPointsById(self):
       self.client.get("/GetCombatPointsById?id=10101")
    @task
    def GetRarityById(self):
       self.client.get("/GetRarityById?id=10101")
    @task
    def GetSoldiersByUn(self):
       self.client.get("/GetSoldiersByUn?un=3")
    @task
    def GetSoldiersByRUCv(self):
       self.client.get("/GetSoldiersByRUCv?Rarity=2&UnlockArena=3&Cvc=1000")